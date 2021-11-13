package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "Please pretend this is a secret message for someone u loved."
	password := "somePassword"

	// bs = byte slice
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}

	bs = bs[:16]

	result, err := enDecode(bs, msg)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result))

	// decode
	result2, err := enDecode(bs, string(result))

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result2))
}

func enDecode(key []byte, input string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	// fmt.Println("debug", block)

	if err != nil {
		return nil, fmt.Errorf("couldn't new cipher %w", err)
	}

	// iv = initialization vector
	iv := make([]byte, aes.BlockSize)

	s 		:= cipher.NewCTR(block, iv)
	buff 	:= &bytes.Buffer{}
	sw 		:= cipher.StreamWriter {
		S: s,
		W: buff,
	}

	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't sw.Write to stream writer %w", err)
	}

	return buff.Bytes(), nil
}
