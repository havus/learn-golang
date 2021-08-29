package web_basic

import (
	"html/template"
	"net/http"
	"testing"
	"net/http/httptest"
	"fmt"
	"io"
	_ "embed"
	"embed"
)

func TestSubjectTemplateText(t *testing.T) {
	server := http.Server{
		Addr: "localhost:3000",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			templateText := `<html><body>{{.}}</body></html>`

			// t, err := template.New("SAMPLE").Parse(templateText)
			// if err != nil {
			// 	panic(err)
			// }
			t := template.Must(template.New("SAMPLE").Parse(templateText))

			t.ExecuteTemplate(writer, "SAMPLE", "Hello World, lorem ipsum...")
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSubjectTemplateFile(t *testing.T) {
	// ============ MANUAL
	// server := http.Server{
	// 	Addr: "localhost:3000",
	// 	Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	// 		t := template.Must(template.ParseFiles("./templates/sample.gohtml"))

	// 		t.ExecuteTemplate(writer, "sample.gohtml", "Hello World, with file lorem ipsum...")
	// 	}),
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }

	// ============ TESTING
	httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	reecorder := httptest.NewRecorder()

	func() {
		t := template.Must(template.ParseFiles("./templates/sample.gohtml"))

		t.ExecuteTemplate(reecorder, "sample.gohtml", "Hello World, with file lorem ipsum...")
	}()

	body, _ := io.ReadAll(reecorder.Result().Body)
	fmt.Println(string(body))
}

func TestSubjectTemplateFolder(t *testing.T) {
	httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	reecorder := httptest.NewRecorder()

	func() {
		t := template.Must(template.ParseGlob("./templates/*.gohtml"))

		t.ExecuteTemplate(reecorder, "sample.gohtml", "lorem ipsum...")
	}()

	body, _ := io.ReadAll(reecorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

type Address struct {
	Street 					string
	Number 					int
	LockdownStatus 	string
}

type Page struct {
	Title string
	Body string
	Address Address
	Hobbies []string
}

func TestSubjectTemplateEmbed(t *testing.T) {
	httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	reecorder := httptest.NewRecorder()

	func() {
		t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

		// t.ExecuteTemplate(reecorder, "name.gohtml", map[string]interface{}{
		// 	"Title": "title",
		// 	"Body": "Wow, this is content!",
		// })
		t.ExecuteTemplate(reecorder, "name.gohtml", Page{
			Title: "title",
			Body: "Wow, this is content!",
			Address: Address{
				Street: "Jakarta",
				Number: 7,
				LockdownStatus: "Urgent",
			},
			Hobbies: []string{
				"Read movie",
				"Write movie",
			},
		})
	}()

	body, _ := io.ReadAll(reecorder.Result().Body)
	fmt.Println(string(body))
}
