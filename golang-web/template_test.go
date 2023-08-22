package golang_web

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed public/layout.html
var layout string

func SimpleHTML(w http.ResponseWriter, r *http.Request){
	templateText := layout
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil{
		panic(err)
	}

	t.ExecuteTemplate(w, "SIMPLE", `Hello Golang`)	
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil{
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", `<div className="container my-5">Hello Golang Web Templates</div>`)
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil{
		panic(err)
	}
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Golang Embed Template")
}

func TestTemplateLayout(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", SimpleHTML)
	mux.HandleFunc("/template", SimpleHTMLFile)
	mux.HandleFunc("/template/glob", TemplateDirectory)
	mux.HandleFunc("/template/embed", TemplateEmbed)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}

func TestTemplate(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateHTML(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct{
	Title string
	Content string
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/layout.gohtml"))
	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title" : "Golang Template Data",
		"Content" : "Golang Web Template Data Example",
	})
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/layout.gohtml"))
	t.ExecuteTemplate(w, "layout.gohtml", Page{
		Title: "Golang Template Data",
		Content: "Golang Web Template Data Struct",
	})
}

func TestTemplateData(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateDataStruct)
	mux.HandleFunc("/map", TemplateDataMap)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}