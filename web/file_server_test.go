package web_basic

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func TestSubjectServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Query().Get("name") != "" {
				http.ServeFile(writer, request, "./src/index.html")
			} else {
				http.ServeFile(writer, request, "./src/not-found.html")
			}
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed src/index.html
var sourceOk string

//go:embed src/not-found.html
var sourceNotFound string

func TestSubjectServeFileWithEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Query().Get("name") != "" {
				fmt.Fprint(writer, sourceOk)
				} else {
				fmt.Fprint(writer, sourceNotFound)
			}
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

