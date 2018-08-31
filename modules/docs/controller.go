package docs

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func alerts(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/alerts.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func avatars(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/avatars.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func buttons(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/buttons.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func cards(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/cards.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func charts(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/charts.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func colors(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/colors.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func form(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/form-components.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func grid(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/grid.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func tags(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/tags.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func typography(c echo.Context) error {
	return c.Render(http.StatusOK, "docs/typography.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
