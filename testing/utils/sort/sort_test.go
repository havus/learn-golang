package sort

import (
	"strconv"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	// Init
	elements 					:= []int{9, 8, 5, 3, 1, 6}
	expectedElements 	:= []int{1, 3, 5, 6, 8, 9}

	timeoutChannel := make(chan bool, 1)
	defer close(timeoutChannel)

	go func() {
		// Execution
		BubbleSortAsc(elements)
		timeoutChannel <- false
	}()

	go func() {
		time.Sleep(1 * time.Second)
		timeoutChannel <- true
	}()

	if <- timeoutChannel {
		t.Error("Bubble sort took a long time")
		return
	}

	// Validation
	for i := 0; i < len(expectedElements); i++ {
		expectedElement := expectedElements[i]

		if elements[i] != expectedElement {
			t.Error("element " + strconv.Itoa(i) + "should be " + strconv.Itoa(expectedElement) + "")
		}
	}
}

// go test -bench=BenchmarBubbleSort 
func BenchmarBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSortAsc([]int{9, 8, 5, 3, 1, 6})
	}
}
