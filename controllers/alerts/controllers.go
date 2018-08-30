package alerts

import (
	"net/http"

	"github.com/labstack/echo"
)

func List(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
