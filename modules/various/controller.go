package various

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mia0x75/dashboard-go/g"
	"github.com/mia0x75/dashboard-go/models/uic"
	"github.com/mia0x75/dashboard-go/utils"
)

var (
	startTime = time.Now()
)

func dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func users(c echo.Context) error {
	return c.Render(http.StatusOK, "users-list.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func currencies(c echo.Context) error {
	return c.Render(http.StatusOK, "crypto-currencies.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func pagination(c echo.Context) error {
	return c.Render(http.StatusOK, "pagination.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func lookup(c echo.Context) error {
	return c.Render(http.StatusOK, "lookup.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func invoice(c echo.Context) error {
	return c.Render(http.StatusOK, "invoice.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func sample(c echo.Context) error {
	return c.Render(http.StatusOK, "sample-cards.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func login(c echo.Context) error {
	if c.Request().Method == "GET" {
		token := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"name": "Dolly!",
			"csrf": token,
		})
	}
	if c.Request().Method == "POST" {
		user := c.FormValue("user")
		password := c.FormValue("password")
		u := &uic.User{
			Login: user,
		}
		g.Con().Uic.Ping()
		if has, err := g.Con().Uic.Get(u); err == nil && has {
			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
				// Create token
				token := jwt.New(jwt.SigningMethodHS256)

				// Set claims
				claims := token.Claims.(jwt.MapClaims)
				claims["id"] = u.Id
				claims["user"] = u.Login
				claims["name"] = u.Name
				claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

				// Generate encoded token and send it as response.
				cipher, err := token.SignedString([]byte(viper.GetString("secret")))
				if err != nil {
					return err
				}

				cookie := new(http.Cookie)
				cookie.Name = "token"
				cookie.Value = cipher
				cookie.Expires = time.Now().Add(24 * time.Hour)
				c.SetCookie(cookie)

				return c.JSON(http.StatusOK, nil)
			}
		} else if err != nil {
			// Unexpected error occured
			fmt.Printf("%v", err)
		}
	}
	return echo.NewHTTPError(http.StatusBadRequest, "Method not allowed.")
}

func register(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func forgot(c echo.Context) error {
	return c.Render(http.StatusOK, "forgot-password.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error400(c echo.Context) error {
	return c.Render(http.StatusOK, "400.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error401(c echo.Context) error {
	return c.Render(http.StatusOK, "401.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error402(c echo.Context) error {
	return c.Render(http.StatusOK, "402.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error403(c echo.Context) error {
	return c.Render(http.StatusOK, "403.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error404(c echo.Context) error {
	return c.Render(http.StatusOK, "404.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error500(c echo.Context) error {
	return c.Render(http.StatusOK, "500.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func error503(c echo.Context) error {
	return c.Render(http.StatusOK, "503.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func about(c echo.Context) error {
	updateSystemStatus()
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name":  "Dolly!",
		"title": "page_title",
		"stats": stats,
	})
}

var stats struct {
	Uptime       string
	NumGoroutine int

	// General statistics.
	MemAllocated string // bytes allocated and still in use
	MemTotal     string // bytes allocated (even if freed)
	MemSys       string // bytes obtained from system (sum of XxxSys below)
	Lookups      uint64 // number of pointer lookups
	MemMallocs   uint64 // number of mallocs
	MemFrees     uint64 // number of frees

	// Main allocation heap statistics.
	HeapAlloc    string // bytes allocated and still in use
	HeapSys      string // bytes obtained from system
	HeapIdle     string // bytes in idle spans
	HeapInuse    string // bytes in non-idle span
	HeapReleased string // bytes released to the OS
	HeapObjects  uint64 // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // bootstrap stacks
	StackSys    string
	MSpanInuse  string // mspan structures
	MSpanSys    string
	MCacheInuse string // mcache structures
	MCacheSys   string
	BuckHashSys string // profiling bucket hash table
	GCSys       string // GC metadata
	OtherSys    string // other system allocations

	// Garbage collector statistics.
	NextGC       string // next run in HeapAlloc time (bytes)
	LastGC       string // last run in absolute time (ns)
	PauseTotalNs string
	PauseNs      string // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
	NumGC        uint32
}

func updateSystemStatus() {
	stats.Uptime = utils.TimeSincePro(startTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	stats.NumGoroutine = runtime.NumGoroutine()

	stats.MemAllocated = utils.FileSize(int64(m.Alloc))
	stats.MemTotal = utils.FileSize(int64(m.TotalAlloc))
	stats.MemSys = utils.FileSize(int64(m.Sys))
	stats.Lookups = m.Lookups
	stats.MemMallocs = m.Mallocs
	stats.MemFrees = m.Frees

	stats.HeapAlloc = utils.FileSize(int64(m.HeapAlloc))
	stats.HeapSys = utils.FileSize(int64(m.HeapSys))
	stats.HeapIdle = utils.FileSize(int64(m.HeapIdle))
	stats.HeapInuse = utils.FileSize(int64(m.HeapInuse))
	stats.HeapReleased = utils.FileSize(int64(m.HeapReleased))
	stats.HeapObjects = m.HeapObjects

	stats.StackInuse = utils.FileSize(int64(m.StackInuse))
	stats.StackSys = utils.FileSize(int64(m.StackSys))
	stats.MSpanInuse = utils.FileSize(int64(m.MSpanInuse))
	stats.MSpanSys = utils.FileSize(int64(m.MSpanSys))
	stats.MCacheInuse = utils.FileSize(int64(m.MCacheInuse))
	stats.MCacheSys = utils.FileSize(int64(m.MCacheSys))
	stats.BuckHashSys = utils.FileSize(int64(m.BuckHashSys))
	stats.GCSys = utils.FileSize(int64(m.GCSys))
	stats.OtherSys = utils.FileSize(int64(m.OtherSys))

	stats.NextGC = utils.FileSize(int64(m.NextGC))
	stats.LastGC = fmt.Sprintf("%.3fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	stats.PauseTotalNs = fmt.Sprintf("%.3fs", float64(m.PauseTotalNs)/1000/1000/1000)
	stats.PauseNs = fmt.Sprintf("%.6fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	stats.NumGC = m.NumGC
}
