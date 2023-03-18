package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"NI/waste"
)

const Version = "1.2"

var (
	FlagCPU                    = flag.Duration("cpu", 0, "Interval of CPU streess (enables CPU stress)")
	FlagCPUduration            = flag.Duration("cpu-d", 2*time.Second, "Min. duration for each CPU stress")
	FlagCPUpercent             = flag.Float64("cpu-p", 100.0, "Each CPU's load percentage")
	FlagCPUcount               = flag.Int("cpu-n", runtime.NumCPU(), "Number of CPU cores to stress")
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
		fmt.Println("Starting", *FlagCPUpercent, "% load of", *FlagCPUcount,"CPUs with interval", *FlagCPU, "of min duration", *FlagCPUduration, "each")
		go waste.CPU(*FlagCPU, *FlagCPUduration, *FlagCPUpercent, *FlagCPUcount)
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
		flag.PrintDefaults()
	} else {
		// fatal error: all goroutines are asleep - deadlock!
		// select {} // fall asleep

		for {
			time.Sleep(24 * time.Hour)
		}
	}
}
