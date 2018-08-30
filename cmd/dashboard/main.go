package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/mia0x75/dashboard-go/controllers/alerts"
	"github.com/mia0x75/dashboard-go/controllers/hosts"
	"github.com/mia0x75/dashboard-go/g"
	"github.com/mia0x75/dashboard-go/hack"
	"github.com/mia0x75/dashboard-go/utils"
)

var (
	path string
)

func init() {
	var err error
	path, err = utils.GetCurrentPath()
	if err != nil {
		fmt.Printf("cannot startup project, error: %s\r\n", err.Error())
		os.Exit(1)
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

	// Static files
	e.Static("/assets", path+"/public/assets")
	e.Static("/demo", path+"/public/demo")
	// Favicon
	e.File("/favicon.ico", path+"/public/favicon.ico")

	r := e.Group("/api")
	// JSON Web Token middleware
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("secret")),
		ContextKey: viper.GetString("jwt.context_key"),
		AuthScheme: viper.GetString("jwt.auth_scheme"),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login.html" || c.Path() == "/register.html" {
				return true
			}
			return false
		},
	}))

	hosts.Routes(e)
	alerts.Routes(e)

	// Startup http service
	addr := fmt.Sprintf("%s:%d", viper.GetString("listen"), viper.GetInt("port"))
	go func() {
		e.StartTLS(addr, fmt.Sprintf("%s/etc/cert.pem", path), fmt.Sprintf("%s/etc/key.pem", path))
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
