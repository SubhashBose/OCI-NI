package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/layou233/neveridle/waste"
)

const Version = "1.2"

var (
	FlagCPU                    = flag.Duration("c", 0, "Interval for CPU load")
	FlagCPUduration            = flag.Duration("d", 2*time.Second, "Min duration for each CPU load")
	FlagCPUpercent             = flag.Float64("p", 100.0, "CPU load percentage")
	FlagCPUcount               = flag.Int("ncpu", runtime.NumCPU(), "Number of CPU cores to load")
	FlagMemory                 = flag.Float64("m", 0, "GiB of memory use")
	FlagNetwork                = flag.Duration("n", 0, "Interval for network speed test")
	FlagNetworkConnectionCount = flag.Int("t", 10, "Set concurrent connections for network speed test")
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
