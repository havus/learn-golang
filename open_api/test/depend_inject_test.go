package test

import (
	"fmt"
	"open_api/depend_inject"
	"testing"
)

func TestDependInjectTest(t *testing.T) {
	dependInject := depend_inject.InitializeService()
	fmt.Println(dependInject.SimpleRepository)
}