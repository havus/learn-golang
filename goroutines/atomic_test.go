package learn_goroutines

import (
	"fmt"
	"testing"
	"sync"
	"sync/atomic"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()

	fmt.Println("Finish looping at", counter)
}
