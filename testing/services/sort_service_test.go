package services

import (
	"strconv"
	"testing"
)

func TestSortService(t *testing.T) {
	// Init
	elements 					:= []int{9, 8, 5, 3, 1, 6}
	expectedElements 	:= []int{1, 3, 5, 6, 8, 9}

	// Execution
	Sort(elements)

	// Validation
	for i := 0; i < len(expectedElements); i++ {
		expectedElement := expectedElements[i]

		if elements[i] != expectedElement {
			t.Error("element " + strconv.Itoa(i) + "should be " + strconv.Itoa(expectedElement) + "")
		}
	}
}
