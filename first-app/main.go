package main

import (
	"fmt"
	greeting "github.com/havus/go-fluffy-waffle"
	"runtime"
)

func main() {
	fmt.Println(greeting.Greet("Sukarno"))
	fmt.Println(runtime.GOOS) // darwin
}