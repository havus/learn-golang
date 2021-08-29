package learn_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestTimer(t *testing.T) {
	// timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	// // time := <- time.NewTimer(3 * time.Second).C
	// time := <- timer.C
	// fmt.Println(time)

	// Second way
	channel := time.After(1 * time.Second)
	fmt.Println(<- channel)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	// running async, require "wait group"
	time.AfterFunc(2 * time.Second, func() {
		defer group.Done()

		fmt.Println("Execute function")
	})

	group.Wait()
}

func TestTicker(t *testing.T) {
	// looping every 1 second
	ticker 	:= time.NewTicker(1 * time.Second)
	done 		:= make(chan bool)

	time.AfterFunc(3 * time.Second, func() {
		ticker.Stop()
		done <- true
	})

	// for tick := range ticker.C {
	// 	fmt.Println(tick)
	// }
	for {
		select {
		case <- done:
			return
		case t := <- ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}

func TestTick(t *testing.T) {
	// looping every 1 second,
	// While Tick is useful for clients that have no need to shut down the Ticker
	channelTick := time.Tick(1 * time.Second)

	for next := range channelTick {
		fmt.Printf("%v %s\n", next, "RUN AGAIN")
	}
}
