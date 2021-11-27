package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	fmt.Println("running the server...")

	http.ListenAndServe(":8080", nil)
}

func getCode(msg string) string {
	h := hmac.New(sha256.New, []byte("i love thursdays when it rains 8723 inches"))
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
}

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey string = "i love thursdays when it rains 8723 inches"

func getJWT(msg string) (string, error) {
	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// ss = signed string
	ss, err := token.SignedString([]byte(myKey))

	if err != nil {
		return "", fmt.Errorf("couldn't SignedString %w", err)
	}

	return ss, nil
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// code := getCode(email)

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "couldn't getJWT", http.StatusInternalServerError)
		return
	}

	// "hash / message digest / digest / hash value" | "what we stored"
	c := http.Cookie{
		Name:  "session",
		// Value: code + "|" + email,
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	// ================= WITHOUT JWT =================
	// isEqual := true
	// xs := strings.SplitN(c.Value, "|", 5)

	// if len(xs) == 2 {
	// 	cCode := xs[0]
	// 	cEmail := xs[1]

	// 	code := getCode(cEmail)

	// 	isEqual = hmac.Equal([]byte(cCode), []byte(code))
	// }
	// ================= WITHOUT JWT =================
	
	
	
	
	// ================= WITH JWT =================
	ss := c.Value

	// ParseWithClaims return Token, err
	// type Token struct {
	// 	Raw       string                 // The raw token.  Populated when you Parse a token
	// 	Method    SigningMethod          // The signing method used or to be used
	// 	Header    map[string]interface{} // The first segment of the token
	// 	Claims    Claims                 // The second segment of the token
	// 	Signature string                 // The third segment of the token.  Populated when you Parse a token
	// 	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
	// }
	afterVerifToken, err := jwt.ParseWithClaims(ss, &myClaims{}, func(beforeVerifToken *jwt.Token) (interface{}, error) {
		isNotSameAlg := beforeVerifToken.Method.Alg() != jwt.SignedMethodHS256.Alg()

		if isNotSameAlg {
			return nil, fmt.Errorf("someone tried to hack changed signing method")
		}
		
		return []byte(myKey), nil
	})

	// if err != nil {
	// 	http.Error(w, "couldn't ParseWithClaims", http.StatusInternalServerError)
	// }

	isEqual := err == nil && afterVerifToken.Valid

	// ================= WITH JWT =================

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
		// ================= WITH JWT =================
		claims, _ := afterVerifToken.Claims.(*myClaims)
		fmt.Println(claims.Email)
		fmt.Println(claims.ExpiresAt)
		// ================= WITH JWT =================
	}


	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>HMAC Example</title>
	</head>
	<body>
		<p>Cookie value: ` + c.Value + `</p>
		<p>` + message + `</p>
		<form action="/submit" method="post">
			<input type="email" name="email" />
			<input type="submit" />
		</form>
	</body>
	</html>`
	io.WriteString(w, html)
}