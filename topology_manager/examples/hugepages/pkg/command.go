package hugepages

// #include "hugepages.c"
import "C"
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

// Generates Heap Pressure using C
func addHeapPressure() {
	defer wg.Done()
	C.generatePressure()
}

// Adds to an existing Map
// func addHeapPressure() {
// 	defer wg.Done()
// 	var storage bytes.Buffer
// 	for i := 0; i < 50000000; i++ {
// 		// 1 Gbits
// 		if storage.Len() >= 1_000_000_000 {
// 			break
// 		}

// 		storage.WriteString("abcdefghijklmnopqrstuvwxyz|")
// 	}
// 	fmt.Printf("[CREATED] Heap Pressure generated is : %.2f MB\n", float64(storage.Len()/1024/1024))
// 	wg.Wait()
// }

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
