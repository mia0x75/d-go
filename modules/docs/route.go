package docs

import (
	"log"
	"net/http"
	"time"

	"github.com/adam-hanna/jwt-auth/jwt"
	"github.com/labstack/echo"
)

// Routes set handler for /docs
func Routes(e *echo.Echo) {
	r := e.Group("/docs")
	var auth jwt.Auth
	err := jwt.New(&auth, jwt.Options{
		SigningMethodString:   "RS256",
		PrivateKeyLocation:    "etc/rsa",     // `$ openssl genrsa -out app.rsa 2048`
		PublicKeyLocation:     "etc/rsa.pub", // `$ openssl rsa -in app.rsa -pubout > app.rsa.pub`
		RefreshTokenValidTime: 72 * time.Hour,
		AuthTokenValidTime:    15 * time.Minute,
		Debug:                 false,
		IsDevEnv:              false,
	})
	if err != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(err)
	}
	auth.SetUnauthorizedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "I pitty the fool who is unauthorized", 401)
		return
	}))
	auth.SetErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "I pitty the fool who has a 500 internal server error", 500)
		return
	}))
	auth.SetRevokeTokenFunction(func(jti string) error {
		return nil
	})
	auth.SetCheckTokenIdFunction(func(jti string) bool {
		return true
	})
	r.Use(echo.WrapMiddleware(auth.Handler))

	r.GET("/index.html", index)
	r.GET("/alerts.html", alerts)
	r.GET("/avatars.html", avatars)
	r.GET("/buttons.html", buttons)
	r.GET("/cards.html", cards)
	r.GET("/charts.html", charts)
	r.GET("/colors.html", colors)
	r.GET("/form-components.html", form)
	r.GET("/grid.html", grid)
	r.GET("/tags.html", tags)
	r.GET("/typography.html", typography)
}
