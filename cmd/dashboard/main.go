package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/mia0x75/dashboard-go/hack"
	"github.com/mia0x75/dashboard-go/utils"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

var (
	lock      *sync.Mutex  = new(sync.Mutex)
	caches    *cache.Cache = cache.New(5*time.Minute, 10*time.Minute)
	path      string
	templates map[string]*template.Template
)

const (
	JWT_ContextKey = "user"
	JWT_AuthScheme = "Token"

	PAGE_SIZE = 15
)

func parseFiles(t *template.Template, filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		if strings.Index(filename, "views/docs") > 0 {
			name = filename[len(path)+1+len("templates/views/"):] //
		}
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func init() {
	var err error
	path, err = utils.GetCurrentPath()
	if err != nil {
		fmt.Printf("cannot startup project, error: %s\r\n", err.Error())
		os.Exit(1)
	}
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := path + "/templates/"
	pages := []string{}
	err = filepath.Walk(templatesDir+"views/", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(path, ".html") {
			if err != nil {
				return err
			}
			if _, err := ioutil.ReadFile(path); err != nil {
				return err
			}
			pages = append(pages, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	partials, err := filepath.Glob(templatesDir + "partials/*.html")
	if err != nil {
		log.Fatal(err)
	}
	includes := append(layouts, partials...)
	// Generate our templates map from our layouts/ and partials/ directories
	for _, page := range pages {
		files := append(includes, page)
		key := page[len(templatesDir+"views/"):len(page)]
		templates[key] = template.Must(parseFiles(nil, files...))
	}
}

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Ensure the template exists in the map.
	key := name
	tmpl, ok := templates[key]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	err := tmpl.ExecuteTemplate(w, key, data) // Raise an error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func main() {
	cfgTmp := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	cfg := *cfgTmp

	fmt.Println(fmt.Sprintf("git commit: %s", hack.Version))
	fmt.Println(fmt.Sprintf("build time: %s", hack.Compile))
	runtime.GOMAXPROCS(runtime.NumCPU())

	viper.AddConfigPath(".")
	viper.AddConfigPath("/")
	viper.AddConfigPath("./etc")
	viper.SetConfigType("json")
	cfg = strings.Replace(cfg, ".json", "", 1)
	viper.SetConfigName(cfg)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 初始化服务器
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = false
	renderer := &Template{
		templates: templates,
	}
	e.Renderer = renderer
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var code = http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if !c.Response().Committed {
			err = c.NoContent(code)
		}
	}

	// 处理中间件
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	// e.Use(middleware.CSRF()) // TODO: 403
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `time:${time_unix} remote_ip:${remote_ip} host:${host} ` +
			`method:${method} uri:${uri} status:${status} bytes_in:${bytes_in} ` +
			`bytes_out:${bytes_out}` + "\n",
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	// 处理静态资源
	e.Static("/assets", path+"/public/assets")
	e.Static("/demo", path+"/public/demo")

	// 处理路由
	e.File("/favicon.ico", path+"/public/favicon.ico")

	r := e.Group("/api")
	// Default ContextKey for JWT is user
	// Retreive context via c.GET("user")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("secret")),
		ContextKey: JWT_ContextKey,
		AuthScheme: JWT_AuthScheme,
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login.html" || c.Path() == "/register.html" {
				return true
			}
			return false
		},
	}))

	Router(e)

	// 启动服务
	addr := fmt.Sprintf("%s:%d", viper.GetString("listen"), viper.GetInt("port"))
	go func() {
		err := e.StartTLS(addr, fmt.Sprintf("%s/etc/cert.pem", path), fmt.Sprintf("%s/etc/key.pem", path))
		if err != nil {
			e.Logger.Info("shutting down the server")
		} else {
			e.Logger.Info(fmt.Sprintf("The server is running on: %s.", addr))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	sc := make(chan os.Signal)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sc
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
	}()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func getOptionalSoleRequestValue(values url.Values, key string, initial string) (string, error) {
	if value, found := values[key]; found {
		if len(value) == 1 {
			if len(value[0]) > 0 {
				return value[0], nil
			}
		}
	} else {
		return initial, nil
	}
	return "", errors.New("Bad Request")
}

func getRequiredSoleRequestValue(values url.Values, key string) (string, error) {
	if value, found := values[key]; found {
		if len(value) == 1 {
			if len(value[0]) > 0 {
				return value[0], nil
			}
		}
	}
	return "", errors.New("Bad Request")
}

func getClaims(ts string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret")), nil
	})
	if err == nil {
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return claims, nil
		} else {
			return nil, errors.New(fmt.Sprintf("string %s is not a valid token.", ts))
		}
	} else {
		return nil, err
	}
}

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/index.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/cards.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "cards.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/charts.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "charts.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/pricing-cards.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "pricing-cards.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/maps.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "maps.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/icons.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "icons.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/store.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "store.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/blog.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "blog.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/carousel.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "carousel.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.GET("/profile.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/login.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/register.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/forgot-password.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "forgot-password.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/400.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "400.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/401.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "401.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/402.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "402.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/403.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "403.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/404.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "404.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/500.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "500.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/503.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "503.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/email.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "email.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/empty.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "empty.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/rtl.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rtl.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.GET("/gallery.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "gallery.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.GET("/form-elements.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "form-elements.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})

	e.GET("/docs/index.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/alerts.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/alerts.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/avatars.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/avatars.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/buttons.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/buttons.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/cards.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/cards.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/charts.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/charts.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/colors.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/colors.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/form-components.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/form-components.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/grid.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/grid.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/tags.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/tags.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/docs/typography.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "docs/typography.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
}
