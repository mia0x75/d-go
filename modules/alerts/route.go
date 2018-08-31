package alerts

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	h := func(echo.Context) error { return nil }

	e.GET("/cards.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return cards
	})
	e.GET("/charts.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return charts
	})
	e.GET("/pricing-cards.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return pricing
	})
	e.GET("/maps.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return maps
	})
	e.GET("/icons.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return icons
	})
	e.GET("/store.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return store
	})
	e.GET("/blog.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return blog
	})
	e.GET("/carousel.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return carousel
	})
	e.GET("/profile.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return profile
	})
	e.GET("/email.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return email
	})
	e.GET("/empty.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return empty
	})
	e.GET("/rtl.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return rtl
	})
	e.GET("/gallery.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return gallery
	})
	e.GET("/form-elements.html", h, func(next echo.HandlerFunc) echo.HandlerFunc {
		return form
	})
}
