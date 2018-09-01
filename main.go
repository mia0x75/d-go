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

	"github.com/dgrijalva/jwt-go"
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

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = viper.GetBool("debug")
	e.Renderer = &g.Template{}
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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `time:${time_unix} remote_ip:${remote_ip} host:${host} ` +
			`method:${method} uri:${uri} status:${status} bytes_in:${bytes_in} ` +
			`bytes_out:${bytes_out}` + "\n",
	}))
	// TODO: 403
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	alerts.Routes(e)
	docs.Routes(e)
	hosts.Routes(e)
	various.Routes(e)

	// Startup http service
	go func() {
		addr := viper.GetString("addr")
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

func getClaims(cipher string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cipher, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret")), nil
	})
	if err == nil {
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return claims, nil
		} else {
			return nil, fmt.Errorf("string: %s is not a valid token, err: %v", cipher, err)
		}
	} else {
		return nil, err
	}
}

func tokenValidation(c echo.Context) bool {
	_, err := getClaims(c.FormValue("token"))
	if err == nil {
		// id := uint(claims["id"].(float64))
		return true
	} else {
		return false
	}
}
