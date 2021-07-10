package web_basic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
	"strings"
	"testing"
)

func SubjectBodyForm(writer http.ResponseWriter, request *http.Request) {
	// ============ MANUAL WAY * parsing first ============
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName 	:= request.PostForm.Get("last_name")
	// ============ MANUAL WAY * parsing first ============

	// firstName := request.PostFormValue("first_name")
	// lastName 	:= request.PostFormValue("last_name")

	// https://github.com/golang/go/blob/master/src/net/http/status.go
	if firstName == "" || lastName == "" {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(writer, "Invalid name!")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Your name is %s %s", firstName, lastName)
	}
}

func TestSubjectBodyForm(t *testing.T) {
	requestBody	:= strings.NewReader("first_name=John&last_name=Doe")
	// request 		:= httptest.NewRequest(http.MethodPost,...
	request 		:= httptest.NewRequest("POST", "http://localhost:3000/subject", requestBody)
	recorder 		:= httptest.NewRecorder()

	// assign header for request
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	SubjectBodyForm(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	fmt.Println("status_code:", response.StatusCode)
	fmt.Println("body:", string(body))
}


