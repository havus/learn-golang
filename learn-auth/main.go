package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	// "github.com/satori/go.uuid"
)

// "github.com/satori/go.uuid" imported but not usedcompilerUnusedImport
// could not import github.com/satori/go.uuid (cannot find package "github.com/satori/go.uuid" in any of
// 	/usr/local/go/src/github.com/satori/go.uuid (from $GOROOT)
// 	/Users/sleekr/Developments/GOLANG/src/github.com/satori/go.uuid (from $GOPATH))

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

// var key []byte = []byte("salt")

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	fmt.Println(bcrypt.DefaultCost)

	pass := "123123123"
	hashPass, err := hashPassword(pass)

	if err != nil {
		panic(err)
	}

	if err := comparePassword(pass, hashPass); err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")

	fmt.Println("SHA512 is from here:", len(sha512.New().Sum(nil)) * 8)

	message := []byte("the message")
	signature, _ := signMessage(message)
	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	// res, _ := checkSign([]byte("The message"), signature)
	res, _ := checkSign(message, signature)
	fmt.Println("is same:", res)
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("error while generate bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password),)

	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	// h := hmac.New(sha512.New, key)
	h := hmac.New(sha512.New, keys[currentKid].key)

	if _, err := h.Write(msg); err != nil {
		return nil, fmt.Errorf("error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSign(msg, sign []byte) (bool, error) {
	newSign, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSign while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSign, sign)

	return same, nil
}

func createToken(c *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.signedString(key)

	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}

	return signedToken, nil
}

type key struct {
	key []byte
	createdAt time.Time
}

func generateNewKey() error {
	newKey := make([]byte, 64)

	if _, err := io.ReadAll(rand.Reader, newKey); err != nil {
		return fmt.Errorf("Error in generateNewKey while generating key: %w", err)
	}

	uid, err := uuid.NewV4()

	if err != nil {
		fmt.Errorf("Error in generateNewKey while generating kid: %w", err)
	}

	currentKid := uid.String()

	keys[currentKid] = key{
		key: newKey,
		createdAt: time.Now(),
	}

	return nil
}

var currentKid = ""
// use database looks good
var keys = map[string]key{}

func parseToken(signedToken string) (*UserClaims, error) {
	// var claims *UserClaims
	// claims := &UserClaims{}
	
	parsedToken, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func (t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil,fmt.Errorf("Invalid key ID")
		}

		k, ok := keys[kid]
		if !ok {
			return nil,fmt.Errorf("Invalid key ID")
		}

		return k, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !parsedToken.Valid() {
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}

	return parsedToken.Claims.(*UserClaims), nil
}