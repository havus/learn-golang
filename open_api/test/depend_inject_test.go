package test

import (
	"fmt"
	"open_api/depend_inject"
	"testing"
)

func TestDependInjectTest(t *testing.T) {
	dependInject, err := depend_inject.InitializeService()
	if err != nil {
		fmt.Println(err)
		fmt.Println(dependInject)
	} else {
		fmt.Println(dependInject.SimpleRepository)
	}
}