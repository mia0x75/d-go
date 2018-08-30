package alerts

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Routes(e *echo.Echo) {
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
