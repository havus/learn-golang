package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Please pretend this is a secret message for someone u loved."
	encoded := encode(msg)

	fmt.Println(encoded)

	decoded, err := decode(encoded)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(decoded)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encodedStr string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(encodedStr)

	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}

	return string(decoded), nil
}