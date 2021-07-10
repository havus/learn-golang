package web_basic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
	"strings"
	"testing"
)

func SubjectHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello world response")
}
func TestHttpHandler(t *testing.T) {
	request 	:= httptest.NewRequest("GET", "http://localhost:3000/subject", nil)
	recorder 	:= httptest.NewRecorder()

	SubjectHandler(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	fmt.Println(string(body))
}



func SubjectQueryParams(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Welcome %s %s", firstName, lastName)
}
func TestQueryParams(t *testing.T) {
	request 	:= httptest.NewRequest("GET", "http://localhost:3000/subject?first_name=John&last_name=Doe", nil)
	recorder 	:= httptest.NewRecorder()

	SubjectQueryParams(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	fmt.Println("body:", string(body))
}



func SubjectMultipleParams(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	keys := query["key"]

	fmt.Fprintf(writer, "Keys are %s", strings.Join(keys, ", "))
}
func TestMultipleParams(t *testing.T) {
	request 	:= httptest.NewRequest("GET", "http://localhost:3000/subject?key=John&key=Doe", nil)
	recorder 	:= httptest.NewRecorder()

	SubjectMultipleParams(recorder, request)

	response 	:= recorder.Result()
	body, _ 	:= io.ReadAll(response.Body)

	fmt.Println("StatusCode:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	fmt.Println("body:", string(body))
}


