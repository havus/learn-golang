package web_basic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
	"strconv"
	"strings"
	"testing"
)

func SubjectCookie(writer http.ResponseWriter, request *http.Request) {
	cookie 				:= new(http.Cookie)
	cookie.Name 	= "is-login"
	cookie.Value 	= "true"
	cookie.Path 	= "/"

	cookie2				:= new(http.Cookie)
	cookie2.Name 	= "second-cookie"
	cookie2.Value	= "second"
	cookie2.Path 	= "/"

	http.SetCookie(writer, cookie)
	http.SetCookie(writer, cookie2)
	
	username 	:= request.PostFormValue("username")
	password 	:= request.PostFormValue("password")

	cookies := request.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie from client:", cookie.Name, cookie.Value)
	}

	if username != "admin" && password == "admin" {
		writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(writer, "Invalid username or password!")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Login status is %s!", strconv.FormatBool(true))
	}
}

func TestSubjectCookie(t *testing.T) {
	requestBody	:= strings.NewReader("first_name=admin&last_name=admin")
	request 		:= httptest.NewRequest("POST", "http://localhost:3000/subject", requestBody)
	recorder 		:= httptest.NewRecorder()

	cookie 				:= new(http.Cookie)
	cookie.Name 	= "from_client"
	cookie.Value 	= "yesyes"
	cookie.Path 	= "/"

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(cookie)

	SubjectCookie(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	cookies := response.Cookies()

	for _, cookie := range cookies {
		fmt.Println("cookie_from_server:", cookie.Name, cookie.Value)
	}

	fmt.Println("body:", string(body))
}

