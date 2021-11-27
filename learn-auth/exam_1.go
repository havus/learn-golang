package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

var db = map[string][]byte{}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("errMsg")

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Exam 1</title>
	</head>
	<body>
		<h1>Error: %s</h1>

		<h1>Register Form</h1>
		<form action="/register" method="POST">
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

	fmt.Fprintf(w, html, errMsg)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errMsg := url.QueryEscape("your method was NOT POST")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		errMsg := url.QueryEscape("error while parsing a body")
		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	email := r.PostForm.Get("e")
	log.Printf(email)

	if email == "" {
		errMsg := url.QueryEscape("your email was blank")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}
	
	password := r.PostForm.Get("p")

	if password == "" {
		errMsg := url.QueryEscape("your password was blank")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := "Failed to GenerateFromPassword"

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	db[email] = bsp

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errMsg := url.QueryEscape("your method was NOT POST")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		errMsg := url.QueryEscape("error while parsing a body")
		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	email := r.PostForm.Get("e")
	log.Printf(email)

	if email == "" {
		errMsg := url.QueryEscape("your email was blank")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}
	
	password := r.PostForm.Get("p")

	if password == "" {
		errMsg := url.QueryEscape("your password was blank")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword(db[email], []byte(password)); err != nil {
		errMsg := url.QueryEscape("your email or your password was wrong!!!")

		http.Redirect(w, r, "/?errMsg=" + errMsg, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
