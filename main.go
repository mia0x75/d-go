package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/mia0x75/dashboard-go/g"
	"github.com/mia0x75/dashboard-go/hack"
	"github.com/mia0x75/dashboard-go/modules/alerts"
	"github.com/mia0x75/dashboard-go/modules/docs"
	"github.com/mia0x75/dashboard-go/modules/hosts"
	"github.com/mia0x75/dashboard-go/modules/various"
)

var (
	UnrestrictedResources = [...]*regexp.Regexp{
		regexp.MustCompile("^/login\\.html$"),
		regexp.MustCompile("^/register\\.html$"),
		regexp.MustCompile("^/forgot-password\\.html$"),
		regexp.MustCompile("^/terms\\.html$"),
		regexp.MustCompile("^/assets/.+$"),
		regexp.MustCompile("^/demo/.+$"),
		regexp.MustCompile("^/favicon.ico$"),
	}
)

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
	err = g.InitDB(viper.GetViper())
	if err != nil {
		log.Fatalf("db conn failed with error %s", err.Error())
	}

	// Initialize echo object
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = false
	renderer := &g.Template{}
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

	e.Static("/assets", "public/assets")
	e.Static("/demo", "public/demo")
	e.File("/favicon.ico", "public/assets/images/favicon.ico")

	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.Rewrite(map[string]string{
		"/":   "/index.html",
		"/*/": "/$1/index.html",
	}))
	e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))
	e.Use(middleware.MethodOverride())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CSRF())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `time:${time_unix} remote_ip:${remote_ip} host:${host} ` +
			`method:${method} uri:${uri} status:${status} bytes_in:${bytes_in} ` +
			`bytes_out:${bytes_out}` + "\n",
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	if !viper.GetBool("debug") {
		e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(viper.GetString("secret")),
			ContextKey: viper.GetString("jwt.context_key"),
			AuthScheme: viper.GetString("jwt.auth_scheme"),
			Skipper: func(c echo.Context) bool {
				// Skip authentication for and signup login requests
				for _, p := range UnrestrictedResources {
					if p.MatchString(c.Path()) {
						return true
					}
				}
				return false
			},
			BeforeFunc: func(c echo.Context) {
				c.Set("Authorized", false)
			},
			SuccessHandler: func(c echo.Context) {
				c.Set("Authorized", true)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				// Redirect to login.html
				c.Redirect(http.StatusSeeOther, "/login.html")
				return nil
			},
		}))
	}

	alerts.Routes(e)
	docs.Routes(e)
	hosts.Routes(e)
	various.Routes(e)

	// Startup http service
	addr := fmt.Sprintf("%s:%d", viper.GetString("listen"), viper.GetInt("port"))
	go func() {
		e.StartTLS(addr, "etc/cert.pem", "etc/key.pem")
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
