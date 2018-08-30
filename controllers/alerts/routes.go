package alerts

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routes(e *echo.Echo) {
	r := e.Group("/alerts")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("cfg.Secret"), // TODO:
		ContextKey: "JWT_ContextKey",     // TODO:
		AuthScheme: "JWT_AuthScheme",     // TODO:
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login.html" || c.Path() == "/register.html" {
				return true
			}
			return false
		},
	}))

	r.GET("/", List)
}
