package main

import (
	"fmt"
	"embed"
	_ "embed"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed havus-logo.png
var havusLogo []byte

// using path matcher
//go:embed files/*.txt
var allFiles embed.FS

func main() {
	// run go build main.go
	// NOT realtime: will store all binary of files
	// add file of remove file will not affected to program
	fmt.Println(version)

	err := ioutil.WriteFile("new_logo.png", havusLogo, fs.ModePerm)
	// if err != nil : err * 
	if (err != nil) {
		panic(err)
	}

	dirEntries, err := allFiles.ReadDir("files")

	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := allFiles.ReadFile("files/" + entry.Name()) 
			fmt.Println("content =", string(file))
		}
	}
}