package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("Server running...")

	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(30 * time.Second)
	fmt.Fprint(w, "Hello world")
}