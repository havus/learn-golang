package learn_goroutines

import (
	"fmt"
	"testing"
	"runtime"
	"sync"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			time.Sleep(1 * time.Second)
			fmt.Println("Running...")
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	// GOMAXPROCS by default similar as cpu count
	// runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines", totalGoroutines)

	group.Wait()
}