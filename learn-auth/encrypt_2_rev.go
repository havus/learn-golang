package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// this is revision/refactored code

func main() {
	msg := "Please pretend this is a secret message for someone u loved."
	password := "somePassword"

	// bs = byte slice
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}

	bs = bs[:16]

	wtr := &bytes.Buffer{}
	encWriter, err := encryptWriter(wtr, bs)

	_, err = io.WriteString(encWriter, msg)
	if err != nil {
		log.Fatalln(err)
	}

	encrypted := wtr.String()
	fmt.Println(encrypted)

	// decode
	result, err := enDecode(bs, encrypted)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result))
}

func enDecode(key []byte, input string) ([]byte, error) {
	block, err := aes.NewCipher(key)

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

func encryptWriter(wtr io.Writer, key []byte) (io.Writer, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, fmt.Errorf("couldn't new cipher %w", err)
	}

	// iv = initialization vector
	iv 		:= make([]byte, aes.BlockSize)
	s 		:= cipher.NewCTR(block, iv)

	return cipher.StreamWriter {
		S: s,
		W: wtr,
	}, nil
}