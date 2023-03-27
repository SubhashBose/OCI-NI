package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"NI/waste"
)

const Version = "1.4"

var (
	FlagCPU                    = flag.Duration("cpu", 0, "Interval of CPU streess (enables CPU stress)")
	FlagCPUduration            = flag.Duration("cpu-d", 2*time.Second, "Min. duration for each CPU stress")
	FlagCPUpercent             = flag.Float64("cpu-p", 100.0, "Each CPU's load percentage to generate")
	FlagCPUcount               = flag.Int("cpu-n", runtime.NumCPU(), "Number of CPU cores to stress")
	FlagCPUglobalmaxperent     = flag.Float64("cpu-m", 100.0, "Max limit of system's total CPU load percent")
	FlagMemory                 = flag.Float64("mem", 0, "GiB of memory to use")
	FlagNetwork                = flag.Duration("net", 0, "Interval for network speed test")
	FlagNetworkConnectionCount = flag.Int("net-c", 10, "Set concurrent connections for network speed test")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Version", Version)
	fmt.Println("Platform:", runtime.GOOS, ",", runtime.GOARCH, ",", runtime.Version())

	flag.Parse()
	nothingEnabled := true

	if *FlagMemory != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting memory consumption of", *FlagMemory, "GiB")
		go waste.Memory(*FlagMemory)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagCPU != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting", *FlagCPUpercent, "% load of", *FlagCPUcount,"CPUs with interval of", *FlagCPU, ", min duration", *FlagCPUduration, "each, and max. system  load", *FlagCPUglobalmaxperen)
		go waste.CPU(*FlagCPU, *FlagCPUduration, *FlagCPUpercent, *FlagCPUcount, *FlagCPUglobalmaxperent)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagNetwork != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting network speed testing with interval", *FlagNetwork)
		go waste.Network(*FlagNetwork, *FlagNetworkConnectionCount)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if nothingEnabled {
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nFor duration or interval flags use time units. E.g., 1h5m10s, 5m, 1h, etc.")
	} else {
		// fatal error: all goroutines are asleep - deadlock!
		// select {} // fall asleep

		for {
			time.Sleep(24 * time.Hour)
		}
	}
}
