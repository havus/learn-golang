package learn_goroutines

import (
	"fmt"
	"testing"
	"sync"
)

func TestMap(t *testing.T) {
	data 	:= &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, "->", value)

		return true
	})
}

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	data.Store(value, value)
}