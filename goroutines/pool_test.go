package learn_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestPool(t *testing.T) {
	// pool := sync.Pool{}
	pool := sync.Pool{
		New: func() interface{} {
			return "Default value"
		},
	}

	pool.Put("Hello")
	pool.Put("World")
	pool.Put("Universe")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Operation finished")
}
