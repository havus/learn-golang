package web_basic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWebServer(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// writer.Write() -> byte params
		_, err := fmt.Fprint(writer, "Hello world")

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	server := http.Server{
		Addr: "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestWebServerWithMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("method:", request.Method)
		fmt.Println("request uri:", request.RequestURI)
		fmt.Fprint(writer, "Hello world")
	})
	mux.HandleFunc("/api/v1/resources", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Some resources :)")
	})
	// with trailing slash will be affected to others endpoints,
	// but not affected to /api/v1/resources
	mux.HandleFunc("/api/v1/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Just V1 :(")
	})

	server := http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}