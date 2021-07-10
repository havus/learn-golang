package web_basic

import (
	_ "embed"
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed src
var source embed.FS

func TestSubjectFileServer(t *testing.T) {
	// using manual import
	// directory 	:= http.Dir("./src")
	// fileServer 	:= http.FileServer(directory)

	// using embed
	directory, _ 	:= fs.Sub(source, "src")
	fileServer 		:= http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	// http://localhost:3000/static/
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

