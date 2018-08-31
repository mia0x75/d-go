package various

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	h := func(echo.Context) error { return nil }

	e.GET("/index.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return dashboard
	})
	e.GET("/about.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return about
	})
	e.GET("/login.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return login
	})
	e.GET("/register.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return register
	})
	e.GET("/forgot-password.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return forgot
	})
	e.GET("/400.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error400
	})
	e.GET("/401.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error401
	})
	e.GET("/402.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error402
	})
	e.GET("/403.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error403
	})
	e.GET("/404.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error404
	})
	e.GET("/500.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error500
	})
	e.GET("/503.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return error503
	})
}
