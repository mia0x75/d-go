package various

import (
	"github.com/labstack/echo"
)

// Routes
func Routes(e *echo.Echo) {
	e.GET("/index.html", dashboard)
	e.GET("/about.html", about) // template test purpose
	e.Any("/login.html", login)
	e.GET("/register.html", register)
	e.GET("/forgot-password.html", forgot)
	e.GET("/400.html", error400)
	e.GET("/401.html", error401)
	e.GET("/402.html", error402)
	e.GET("/403.html", error403)
	e.GET("/404.html", error404)
	e.GET("/500.html", error500)
	e.GET("/503.html", error503)

	// e.GET("/captcha", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}
