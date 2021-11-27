package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample-file.txt")
	if err != nil {
		log.Fatalln("log fatal:", err)
	}

	defer f.Close()

	h := sha256.New() // return hash.Hash

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("couldn't io.Copy", err)
	}

	fmt.Printf("here's the type BEFORE Sum: %T\n", h)
	fmt.Printf("%v\n", h)

	sb := h.Sum(nil) // []byte
	fmt.Printf("here's the type AFTER Sum: %T\n", sb)
	fmt.Printf("%x\n", sb)

	sb = h.Sum(nil) // []byte
	fmt.Printf("here's the type AFTER SECOND Sum: %T\n", sb)
	fmt.Printf("%x\n", sb)

	sb = h.Sum(sb) // []byte
	fmt.Printf("here's the type AFTER THIRD Sum with sb: %T\n", sb)
	fmt.Printf("%x\n", sb)
}