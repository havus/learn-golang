package embed_learn

import (
	"fmt"
	"testing"
	"embed"
	_ "embed"
	// "io/fs"
	// "io/ioutil"
)

//go:embed version.txt
var version string

//go:embed havus-logo.png
var havusLogo []byte

//go:embed files/file_1.txt
//go:embed files/file_2.txt
//go:embed files/file_3.txt
//go:embed files2/test.txt
var files embed.FS

// using path matcher
//go:embed files/*.txt
var allFiles embed.FS

func TestString(t *testing.T) {
	fmt.Println(version)

	// err := ioutil.WriteFile("new_logo.png", havusLogo, fs.ModePerm)
	// if (err != nil) {
	// 	panic(err)
	// }

	// file_1, _ := files.ReadFile("files/file_1.txt")
	// file_2, _ := files.ReadFile("files/file_2.txt")
	// file_3, _ := files.ReadFile("files/file_3.txt")
	// file_test, _ := files.ReadFile("files2/test.txt")
	// fmt.Println(string(file_1))
	// fmt.Println(string(file_2))
	// fmt.Println(string(file_3))
	// fmt.Println(string(file_test))

	// PATH MATCHER
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
