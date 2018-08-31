package various

import (
	"net/http"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	h := func(echo.Context) error { return nil }

	e.GET("/index.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return Dashboard
	})
	e.GET("/about.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return About
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
}
