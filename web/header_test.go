package web_basic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
	"testing"
)

func SubjectHeader(writer http.ResponseWriter, request *http.Request) {
	// assign header for response
	writer.Header().Add("x-powered-by", "anon_user")
	contentType := request.Header.Get("content-type")

	fmt.Fprintf(writer, "Content type is %s", contentType)
}
func TestHeader(t *testing.T) {
	request 	:= httptest.NewRequest("GET", "http://localhost:3000/subject?key=John&key=Doe", nil)
	recorder 	:= httptest.NewRecorder()

	// assign header for request
	request.Header.Add("Content-Type", "application/json")

	SubjectHeader(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	responseHeader := response.Header.Get("X-Powered-By")
	fmt.Println("X-Powered-By:", string(responseHeader))
	fmt.Println("body:", string(body))
}


