package utils

import (
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// MiddlewareAuth
func MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthenticated, err := func() (bool, error) {
			if !viper.GetBool("debug") {
				cookie, err := c.Cookie("token")
				if err != nil {
					return false, nil
				}
				if _, err := GetClaims(cookie.Value); err != nil {
					return false, err
				}
			}
			return true, nil
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

func GetClaims(s string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret")), nil
	})
	if err == nil {
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return claims, nil
		} else {
			return nil, errors.New(fmt.Sprintf("string %s is not a valid token", s))
		}
	} else {
		return nil, err
	}
}
