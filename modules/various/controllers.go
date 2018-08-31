package various

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo"
	"github.com/mia0x75/dashboard-go/utils"
)

var (
	startTime = time.Now()
)

func Dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func About(c echo.Context) error {
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
