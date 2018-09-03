package utils

import (
	"net/http"

	"github.com/labstack/echo"
)

// MiddlewareAuth
func MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthenticated, err := func() (bool, error) {
			// TODO:
			return false, nil
		}()
		if err != nil {
			isAuthenticated = false
		}
		if !isAuthenticated {
			return c.Redirect(http.StatusSeeOther, "/login.html")
		}
		return next(c)
	}
}
