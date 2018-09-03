package alerts

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/cards.html", cards)
	e.GET("/charts.html", charts)
	e.GET("/pricing-cards.html", pricing)
	e.GET("/maps.html", maps)
	e.GET("/icons.html", icons)
	e.GET("/store.html", store)
	e.GET("/blog.html", blog)
	e.GET("/carousel.html", carousel)
	e.GET("/profile.html", profile)
	e.GET("/email.html", email)
	e.GET("/empty.html", empty)
	e.GET("/rtl.html", rtl)
	e.GET("/gallery.html", gallery)
	e.GET("/form-elements.html", form)
}
