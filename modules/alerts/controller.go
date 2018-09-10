package alerts

import (
	"net/http"

	"github.com/labstack/echo"
)

func cards(c echo.Context) error {
	return c.Render(http.StatusOK, "cards.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func charts(c echo.Context) error {
	return c.Render(http.StatusOK, "charts.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func pricing(c echo.Context) error {
	return c.Render(http.StatusOK, "pricing-cards.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func maps(c echo.Context) error {
	return c.Render(http.StatusOK, "maps.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func icons(c echo.Context) error {
	return c.Render(http.StatusOK, "icons.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func store(c echo.Context) error {
	return c.Render(http.StatusOK, "store.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func blog(c echo.Context) error {
	return c.Render(http.StatusOK, "blog.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func carousel(c echo.Context) error {
	return c.Render(http.StatusOK, "carousel.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func profile(c echo.Context) error {
	return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func email(c echo.Context) error {
	return c.Render(http.StatusOK, "email.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func empty(c echo.Context) error {
	return c.Render(http.StatusOK, "empty.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func gallery(c echo.Context) error {
	return c.Render(http.StatusOK, "gallery.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func form(c echo.Context) error {
	return c.Render(http.StatusOK, "form-elements.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
