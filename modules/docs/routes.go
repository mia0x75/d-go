package docs

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	h := func(echo.Context) error { return nil }
	r := e.Group("/docs")

	r.GET("/index.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return index
	})
	r.GET("/alerts.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return alerts
	})
	r.GET("/avatars.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return avatars
	})
	r.GET("/buttons.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return buttons
	})
	r.GET("/cards.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return cards
	})
	r.GET("/charts.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return charts
	})
	r.GET("/colors.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return colors
	})
	r.GET("/form-components.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return form
	})
	r.GET("/grid.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return grid
	})
	r.GET("/tags.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return tags
	})
	r.GET("/typography.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return typography
	})
}
