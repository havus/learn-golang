package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Password 	[]byte
	FirstName	string
}

type customClaims struct {
	jwt.StandardClaims
	SID 			string
}

var db 				= map[string]user{}
var sessions 	= map[string]string{}
var key 			= []byte("secret_key")

func main() {
	fmt.Println(createToken("test"))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sid")
	if err != nil {
		cookie = &http.Cookie{
			Name: "sid",
			Value: "",
		}
	}

	sidCookie := cookie.Value
	// s, err := parseToken(sidCookie)
	sID, err := parseToken(sidCookie)

	if err != nil {
		log.Println("index parse token err", err)
	}

	var email string
	if sID != "" {
		email = sessions[sID]
	}

	var firstname string
	if currentUser, ok := db[email]; ok {
		firstname = currentUser.FirstName
	}


	msg := r.FormValue("msg")

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Exam 1</title>
	</head>
	<body>
		<h1>Your e-mail: %s</h1>
		<h1>Message: %s</h1>
		<h1>FirstName: %s</h1>

		<h1>Register Form</h1>
		<form action="/register" method="POST">
			<label for="firstname">First Name<label>
			<input type="text" name="firstname" id="firstname">
			<input type="email" name="e">
			<input type="password" name="p">
			<input type="submit">
		</form>

		<h1>Login Form</h1>
		<form action="/login" method="POST">
			<input type="email" name="e">
			<input type="password" name="p">
			<input type="submit">
		</form>
	</body>
	</html>`

	fmt.Fprintf(w, html, email, msg, firstname)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := url.QueryEscape("your method was NOT POST")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		msg := url.QueryEscape("error while parsing a body")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	firstname := r.PostForm.Get("firstname")
	log.Printf(firstname)

	if firstname == "" {
		msg := url.QueryEscape("your firstname was blank")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	email := r.PostForm.Get("e")
	log.Printf(email)

	if email == "" {
		msg := url.QueryEscape("your email was blank")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}
	
	password := r.PostForm.Get("p")

	if password == "" {
		msg := url.QueryEscape("your password was blank")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		msg := "Failed to GenerateFromPassword"

		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// db[email] = bsp
	db[email] = user{
		Password: bsp,
		FirstName: firstname,
	}
	fmt.Println(db)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := url.QueryEscape("your method was NOT POST")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		msg := url.QueryEscape("error while parsing a body")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	email := r.PostForm.Get("e")
	log.Printf(email)

	if email == "" {
		msg := url.QueryEscape("your email was blank")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	password := r.PostForm.Get("p")

	if password == "" {
		msg := url.QueryEscape("your password was blank")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword(db[email].Password, []byte(password)); err != nil {
		msg := url.QueryEscape("your email or your password was wrong!!!")

		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	sUUID := uuid.New().String()
	sessions[sUUID] = email
	token, err := createToken(sUUID)
	if err != nil {
		log.Printf("error in loginHandler while create token:", err)

		msg := url.QueryEscape("Our server did not get your token for lunch, please take a break now!")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	cookie := http.Cookie{
		Name: 	"sid",
		Value: 	token,
	}

	http.SetCookie(w, &cookie)
	
	msg := url.QueryEscape("You logged in! :)")
	http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
}

func createToken(sid string) (string, error) {
	cc := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
		SID: sid,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	ss, err := jwtToken.SignedString(key)

	if err != nil {
		return "", fmt.Errorf("couldn't sign jwt in createToken %w", err)
	}
	return ss, nil

	// // func New(h func() hash.Hash, key []byte) hash.Hash
	// mac := hmac.New(sha256.New, key)
	// mac.Write([]byte(sid))

	// // signedMAC := mac.Sum(nil)

	// // to base64
	// signedMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// return string(signedMAC) + "|" + sid
}

func parseToken(ss string) (string, error) {
	token, err := jwt.ParseWithClaims(ss, &customClaims{}, func (t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("ParseWithClaims different alg used")
		}

		return key, nil
	})

	if err != nil {
		return "", fmt.Errorf("couldn't ParsewithClaims in parseToken %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token in parseToken %w", err)
	}

	return token.Claims.(*customClaims).SID, nil

	// xs 	:= strings.SplitN(ss, "|", 2)
	// b64 := xs[0]

	// fmt.Println(xs, ss, "debug")
	// if len(xs) != 2 {
	// 	return "", fmt.Errorf("Error format while parseToken")
	// }

	// xb, err := base64.StdEncoding.DecodeString(b64) // []byte{}, err

	// if err != nil {
	// 	return "", fmt.Errorf("Error in parseToken %w", err)
	// }

	// sid := xs[1]
	// originMAC := hmac.New(sha256.New, key)
	// originMAC.Write([]byte(xs[1]))

	// isEqual := hmac.Equal(xb, originMAC.Sum(nil))

	// if !isEqual {
	// 	return "", fmt.Errorf("Error in parseToken hmac.Equal")
	// }

	// return sid, nil
}