package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
)

type user struct {
	Password 	[]byte
	FirstName	string
}

type customClaims struct {
	jwt.StandardClaims
	SID 			string
}

var db 								= map[string]user{}
var sessions 					= map[string]string{}
var key 							= []byte("secret_key")
var oauthExp 					= map[string]time.Time{} // { uuid: expiration_time }
var oauthConn 				= map[string]string{} // { uid from oauth provider: user_id/email }
var amazonOauthConfig = oauth2.Config{
	ClientID:     "amzn1.application-oa2-client.fbf31c80acc94a3b89dd45f2dd438d2d",
	ClientSecret: "64b0ee3063625b02ad3fd5e8f2b72a3036670bf48c37365e8d2fd816a7667e96",
	Endpoint:     amazon.Endpoint,
	RedirectURL: 	"http://localhost:3000/oauth/amazon/receive",
	Scopes: 			[]string{"profile"},
}

func main() {
	fmt.Println(createToken("test"))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	// https://na.account.amazon.com/ap/oa?arb=b822e9fc-6e37-4e52-aa3e-5f41b9b68aef
	http.HandleFunc("/oauth/amazon/login", loginOauthHandler)
	http.HandleFunc("/oauth/amazon/receive", receiveOauthHandler)
	http.HandleFunc("/oauth/amazon/register", amazonRegister)
	http.HandleFunc("/partial-register", partialRegister)
	http.HandleFunc("/logout", logoutHandler)

	http.ListenAndServe(":3000", nil)
}

func amazonRegister(w http.ResponseWriter, r *http.Request) {
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

	oauthID 	:= r.PostForm.Get("oauthID")
	firstname	:= r.PostForm.Get("firstname")
	email			:= r.PostForm.Get("email")

	if email == "" {
		msg := url.QueryEscape("your email was blank")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if oauthID == "" {
		msg := url.QueryEscape("the oauthID was blank")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if firstname == "" {
		msg := url.QueryEscape("your name was blank")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	amazonID, err := parseToken(oauthID)
	if err != nil {
		msg := url.QueryEscape("parseToken at amazonRegister" + err.Error())
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	db[email] = user{FirstName: firstname}
	oauthConn[amazonID] = email

	err = createSession(email, w)
	if err != nil {
		log.Printf("error in amazonRegister while create session:", err)
		msg := url.QueryEscape("Our server did not get your session for lunch, please take a break now!")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func partialRegister(w http.ResponseWriter, r *http.Request) {
	// uv.Add("tokenEscape", tokenEscape)
	// uv.Add("name", aur.Name)
	// uv.Add("email", aur.Email)
	tokenEscape := r.FormValue("tokenEscape")
	name 				:= r.FormValue("name")
	email 			:= r.FormValue("email")

	if tokenEscape == "" {
		msg := url.QueryEscape("couldn't get token from partial register")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}
	
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Amazon Register</title>
	</head>
	<body>
		<form action="/oauth/amazon/register" method="post">
			<label for="firstname">First Name</label>
			<input name="firstname" type="text" id="firstname" value="%s">

			<label for="email">Email</label>
			<input name="email" ype="email" id="email" value="%s">

			<input type="hidden" value="%s" name="oauthID">

			<input type="submit" value="Continue Register with Amazon">
		</form>
	</body>
	</html>`

	fmt.Fprintf(w, html, name, email, tokenEscape)
}

func loginOauthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := uuid.New().String()
	oauthExp[id] = time.Now().Add(time.Hour)

	redirectUrl := amazonOauthConfig.AuthCodeURL(id)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

func receiveOauthHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:3000/oauth/amazon/receive?code=ANhzSyopPPqtwjZLljHs&scope=profile&state=d3067e08-ba64-4fba-89d7-99dc51a3832f
	state := r.FormValue("state")
	code 	:= r.FormValue("code")

	if state == "" {
		msg := url.QueryEscape("error while login with amazon [state empty]")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	if code == "" {
		msg := url.QueryEscape("error while login with amazon [code empty]")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	expiredTime := oauthExp[state]

	if time.Now().After(expiredTime) {
		msg := url.QueryEscape("session login with amazon was expired")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	token, err := amazonOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		msg := url.QueryEscape("failed exchange the code while login with amazon" + err.Error())
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	tokenSource := amazonOauthConfig.TokenSource(r.Context(), token)
	client 			:= oauth2.NewClient(r.Context(), tokenSource)

	res, err := client.Get("https://api.amazon.com/user/profile")
	defer res.Body.Close()

	if err != nil {
		msg := url.QueryEscape("failed to get api.amazon.com/user/profile" + err.Error())
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	// bs, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	msg := url.QueryEscape("failed to read the response" + err.Error())
	// 	http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
	// 	return
	// }

	if res.StatusCode < 200 || res.StatusCode > 299 {
		msg := url.QueryEscape("response is not 200 status code, but" + string(res.StatusCode))
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	// fmt.Println(string(bs))
	// io.WriteString(w, string(bs))

	// {
	// 	"user_id": "amzn1.account.AGTWMU2PVJ4Q3O25W22VTAV5JFYA",
	// 	"name": "Hafidz mahrus",
	// 	"email": "imhavus@gmail.com"
	// }
	type amazonUserResponse struct {
		UserID 	string `json:"user_id"`
		Name 		string `json:"name"`
		Email		string `json:"email"`
	}

	var aur amazonUserResponse

	err = json.NewDecoder(res.Body).Decode(&aur)
	if err != nil {
		msg := url.QueryEscape("can NOT decode json response" + err.Error())
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	userID := aur.UserID

	email, ok := oauthConn[userID]
	if !ok {
		// register at our site
		// for k, _ := range db {
		// 	break
		// }
		token, err := createToken(userID)
		if err != nil {
			log.Printf("error in receiveOauthHandler while create token:", err)
			msg := url.QueryEscape("Our server did not get your token for lunch, please take a break now!")
			http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
			return
		}

		tokenEscape := url.QueryEscape(token)
		uv := url.Values{}
		uv.Add("tokenEscape", tokenEscape)
		uv.Add("name", aur.Name)
		uv.Add("email", aur.Email)
		http.Redirect(w, r, "/partial-register?" + uv.Encode(), http.StatusSeeOther)
		// email = "test@example.com"
	}

	err = createSession(email, w)

	if err != nil {
		log.Printf("error in receiveOauthHandler while create session:", err)
		msg := url.QueryEscape("Our server did not get your session for lunch, please take a break now!")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}

	msg := url.QueryEscape("You logged in! :)")
	http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
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

		<form action="/logout" method="POST">
			<input type="submit" value="Logout">
		</form>

		<br>
		<h1>Login with Amazon</h1>
		<form action="/oauth/amazon/login" method="POST">
			<input type="submit" value="Login with Amazon">
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

	err := createSession(email, w)

	if err != nil {
		log.Printf("error in loginHandler while create token:", err)

		msg := url.QueryEscape("Our server did not get your token for lunch, please take a break now!")
		http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
		return
	}
	
	msg := url.QueryEscape("You logged in! :)")
	http.Redirect(w, r, "/?msg=" + msg, http.StatusSeeOther)
}

func createSession(email string, w http.ResponseWriter) error {
	sUUID := uuid.New().String()
	sessions[sUUID] = email
	token, err := createToken(sUUID)
	if err != nil {
		return fmt.Errorf("couldn't create token in createToken %w", err)
	}

	cookie := http.Cookie{
		Name: 	"sid",
		Value: 	token,
		Path: 	"/",
	}

	http.SetCookie(w, &cookie)

	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

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

	delete(sessions, sID)

	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
}