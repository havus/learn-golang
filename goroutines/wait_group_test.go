package learn_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Finish all go routines")
}

// SYNC ONCE

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once 		:= sync.Once{}
	group 	:= sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter in", counter)
}