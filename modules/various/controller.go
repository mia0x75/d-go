package various

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
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
	"github.com/toolkits/nux"
)

var (
	startTime = time.Now()
)

type memInfo struct {
	Buffers   string
	Cached    string
	MemTotal  string
	MemFree   string
	SwapTotal string
	SwapUsed  string
	SwapFree  string
}

var machineStats struct {
	Procs   int
	Uptime  string
	MemInfo *memInfo
	Load    *nux.Loadavg
}

func init() {
	// TODO：
	go func() {
		cmd := exec.Command("inxi", "-S", "-M", "-C", "-D", "-N", "-s", "-I", "--color=0")
		/*
			System:    Host: office Kernel: 4.17.19-200.fc28.x86_64 x86_64 bits: 64 Desktop: LXDE 0.9.3
			           Distro: Fedora release 28 (Twenty Eight)
			Machine:   Type: Desktop System: LENOVO product: 90G0CTO1WW v: QiTianM410-N000 serial: <root required>
			           Mobo: LENOVO model: 3102 v: SDK0L77767 WIN 3423500608959 serial: <root required>
			           UEFI [Legacy]: LENOVO v: M16KT47A date: 02/06/2018
			CPU:       Topology: Dual Core model: Intel Core i3-7100 bits: 64 type: MT MCP L2 cache: 3072 KiB
			           Speed: 800 MHz min/max: 800/3900 MHz Core speeds (MHz): 1: 800 2: 800 3: 800 4: 800
			Network:   Device-1: Realtek RTL8111/8168/8411 PCI Express Gigabit Ethernet driver: r8169
			Drives:    Local Storage: total: 1.05 TiB used: 101.41 GiB (9.4%)
			           ID-1: /dev/sda vendor: Kingston model: SA400S37120G size: 111.79 GiB
			           ID-2: /dev/sdb vendor: Toshiba model: DT01ACA100 LENOVO size: 931.51 GiB
			           ID-3: /dev/sdc type: USB vendor: SanDisk model: Cruzer Fit size: 29.82 GiB
			Sensors:   System Temperatures: cpu: 29.8 C mobo: 27.8 C
			           Fan Speeds (RPM): N/A
			Info:      Processes: 280 Uptime: 20d 23h 26m Memory: 15.57 GiB used: 6.86 GiB (44.0%) Shell: bash inxi: 3.0.21
		*/
		var inxi bytes.Buffer
		cmd.Stdout = &inxi
		err := cmd.Run()
		if err == nil {
		}
	}()
}

func dashboard(c echo.Context) error {
	UpdateServiceStatus()
	d, h, m, _ := nux.SystemUptime()
	mem, _ := nux.MemInfo()
	load, _ := nux.LoadAvg()

	machineStats.Load = load
	machineStats.MemInfo = &memInfo{
		Buffers:   utils.FileSize(int64(mem.Buffers)),
		Cached:    utils.FileSize(int64(mem.Cached)),
		MemTotal:  utils.FileSize(int64(mem.MemTotal)),
		MemFree:   utils.FileSize(int64(mem.MemFree)),
		SwapTotal: utils.FileSize(int64(mem.SwapTotal)),
		SwapUsed:  utils.FileSize(int64(mem.SwapUsed)),
		SwapFree:  utils.FileSize(int64(mem.SwapFree)),
	}
	machineStats.Uptime = fmt.Sprintf("%d 天, %d 小时, %d 分钟", d, h, m)

	name := "Dolly!"
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name":         name,
		"svcStats":     svcStats,
		"machineStats": machineStats,
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
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name":  "Dolly!",
		"title": "page_title",
	})
}

var svcStats struct {
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

func UpdateServiceStatus() {
	svcStats.Uptime = utils.TimeSincePro(startTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	svcStats.NumGoroutine = runtime.NumGoroutine()

	svcStats.MemAllocated = utils.FileSize(int64(m.Alloc))
	svcStats.MemTotal = utils.FileSize(int64(m.TotalAlloc))
	svcStats.MemSys = utils.FileSize(int64(m.Sys))
	svcStats.Lookups = m.Lookups
	svcStats.MemMallocs = m.Mallocs
	svcStats.MemFrees = m.Frees

	svcStats.HeapAlloc = utils.FileSize(int64(m.HeapAlloc))
	svcStats.HeapSys = utils.FileSize(int64(m.HeapSys))
	svcStats.HeapIdle = utils.FileSize(int64(m.HeapIdle))
	svcStats.HeapInuse = utils.FileSize(int64(m.HeapInuse))
	svcStats.HeapReleased = utils.FileSize(int64(m.HeapReleased))
	svcStats.HeapObjects = m.HeapObjects

	svcStats.StackInuse = utils.FileSize(int64(m.StackInuse))
	svcStats.StackSys = utils.FileSize(int64(m.StackSys))
	svcStats.MSpanInuse = utils.FileSize(int64(m.MSpanInuse))
	svcStats.MSpanSys = utils.FileSize(int64(m.MSpanSys))
	svcStats.MCacheInuse = utils.FileSize(int64(m.MCacheInuse))
	svcStats.MCacheSys = utils.FileSize(int64(m.MCacheSys))
	svcStats.BuckHashSys = utils.FileSize(int64(m.BuckHashSys))
	svcStats.GCSys = utils.FileSize(int64(m.GCSys))
	svcStats.OtherSys = utils.FileSize(int64(m.OtherSys))

	svcStats.NextGC = utils.FileSize(int64(m.NextGC))
	svcStats.LastGC = fmt.Sprintf("%.3fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	svcStats.PauseTotalNs = fmt.Sprintf("%.3fs", float64(m.PauseTotalNs)/1000/1000/1000)
	svcStats.PauseNs = fmt.Sprintf("%.6fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	svcStats.NumGC = m.NumGC
}
