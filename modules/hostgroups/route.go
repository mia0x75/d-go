package hostgroups

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	r := e.Group("/hostgroups")

	r.GET("/", List)
}
