package hugepages

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

// Wait for Shutdown (basically make this app play friendly in a container)
func waitForShutdown() {
	defer wg.Done()
	var sigs = []os.Signal{os.Interrupt, syscall.SIGTERM}
	c := make(chan os.Signal, 2)
	signal.Notify(c, sigs...)
	go func() {
		<-c
		os.Exit(0)
	}()
}

// Report Memory Stats in a 1M Loop
func reportMemStats() {
	defer wg.Done()
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		fmt.Printf("Sys: %.2f MB\n", float64(m.Sys)/1024/1024)
		fmt.Printf("Heap Alloc: %.2f MB\n", float64(m.HeapAlloc)/1024/1024)
		time.Sleep(60 * time.Second)
	}
}

// Adds to an existing Map
func addHeapPressure() {
	defer wg.Done()
	var storage bytes.Buffer
	var m runtime.MemStats
	for i := 0; i < 50000000; i++ {
		if i%100 == 0 {
			// The call to ReadMemStats is slow.
			// The conditional check guards against too many slow checks.
			runtime.ReadMemStats(&m)
			pressure := float64(m.HeapAlloc) / 1024 / 1024
			if pressure >= 1000 {
				fmt.Printf("Heap Pressure is : %.2f MB\n", pressure)
				break
			}
		}

		storage.WriteString("abcdefghijklmnopqrstuvwxyz|")
	}
	fmt.Println("Done... creating Memory Pressure ", storage.Len())
	wg.Wait()
}

// The command waits for signals to shutdown.
// The code orchestrates the Report and adding of HEAP Pressure
func Command() {
	go waitForShutdown()
	wg.Add(1)
	go reportMemStats()
	wg.Add(1)
	go addHeapPressure()
	wg.Add(1)
	wg.Wait()
}
