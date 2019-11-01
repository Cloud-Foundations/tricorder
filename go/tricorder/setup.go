package tricorder

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Cloud-Foundations/tricorder/go/tricorder/units"
)

func getProgramArgs() string {
	return strings.Join(os.Args[1:], "|")
}

func initDefaultMetrics() {
	programArgs := getProgramArgs()
	var totalMemory uint64
	memStatsGroup := NewGroup()
	memStatsGroup.RegisterUpdateFunc(func() time.Time {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		if memStats.Sys >= memStats.HeapReleased {
			totalMemory = memStats.Sys - memStats.HeapReleased
		}
		return time.Now()
	})
	RegisterMetricInGroup(
		"/proc/memory/total",
		&totalMemory,
		memStatsGroup,
		units.Byte,
		"System memory currently allocated to process")
	registerRusage()
	goVersion := runtime.Version()
	var numGoroutines int
	runtimeGroup := NewGroup()
	runtimeGroup.RegisterUpdateFunc(func() time.Time {
		numGoroutines = runtime.NumGoroutine()
		return time.Now()
	})
	RegisterMetricInGroup(
		"/proc/go/num-goroutines",
		&numGoroutines,
		runtimeGroup,
		units.None,
		"Number of goroutines")
	RegisterMetric(
		"/proc/go/version",
		&goVersion,
		units.None,
		"Version of Go runtime")
	if countOpenFileDescriptors() >= 0 {
		RegisterMetric(
			"/proc/io/num-open-file-descriptors",
			countOpenFileDescriptors,
			units.None,
			"Number of open file descriptors")
	}
	RegisterMetric("/proc/name", &os.Args[0], units.None, "Program name")
	RegisterMetric("/proc/args", &programArgs, units.None, "Program args")
	RegisterMetric("/proc/start-time", &appStartTime, units.None,
		"Program start time")
}

func init() {
	initDefaultMetrics()
	initHttpFramework()
	initHtmlHandlers()
	initJsonHandlers()
	initRpcHandlers()
}

func countOpenFileDescriptors() int {
	fdDir, err := os.Open("/proc/self/fd")
	if err != nil {
		return -1
	}
	defer fdDir.Close()
	if dirNames, err := fdDir.Readdirnames(0); err != nil {
		return -1
	} else {
		return len(dirNames)
	}
}
