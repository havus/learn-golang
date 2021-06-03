package main

import (
	"flag"
	"fmt"
)

// go run flag.go -host="hello world" -key=root
func main() {
	var host *string = flag.String("host", "defaultHost", "Just a description")
	*key = flag.String("key", "defaultKey", "Insert ur key")
	// var key *string = flag.String("key", "defaultKey", "Insert ur key")
	// host := flag.String("hostname", "defaultValue", "Just a description")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*key)
}