package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var rsaSampleSecret []byte("THIS_IS_SECRET")

func generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(rsaSampleSecret)
}

func main() {
	token, _ := generateJWT()
}