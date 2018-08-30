package utils

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

func getOptionalSoleRequestValue(values url.Values, key string, initial string) (string, error) {
	if value, found := values[key]; found {
		if len(value) == 1 {
			if len(value[0]) > 0 {
				return value[0], nil
			}
		}
	} else {
		return initial, nil
	}
	return "", errors.New("Bad Request")
}

func getRequiredSoleRequestValue(values url.Values, key string) (string, error) {
	if value, found := values[key]; found {
		if len(value) == 1 {
			if len(value[0]) > 0 {
				return value[0], nil
			}
		}
	}
	return "", errors.New("Bad Request")
}

func getClaims(ts string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret")), nil
	})
	if err == nil {
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return claims, nil
		} else {
			return nil, errors.New(fmt.Sprintf("string %s is not a valid token.", ts))
		}
	} else {
		return nil, err
	}
}
