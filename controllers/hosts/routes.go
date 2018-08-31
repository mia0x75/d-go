package hosts

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	r := e.Group("/hosts")

	r.GET("/", List)
}
