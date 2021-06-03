package main

import (
	"os"
	"fmt"
)

// go run os.go coba argumen
func main() {
	args := os.Args
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])

	hostname, _ := os.Hostname()
	fmt.Println(hostname)
}