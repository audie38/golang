package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request){
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Template Caching")
}

type PageContentData struct{
	Title string
	Body interface{}
}

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request){
	data := PageContentData{
		Title: "Go-Lang Auto Escape",
		Body: `<div class="container my-5 text-center"><p>Go-Lang Auto Escape</p></div>`,
	}

	myTemplates.ExecuteTemplate(w, "post", data)
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request){
	data := PageContentData{
		Title: "Go-Lang Auto Escape",
		Body: template.HTML( `<div class="container my-5 text-center"><p>Go-Lang Auto Escape</p></div>`),
	}

	myTemplates.ExecuteTemplate(w, "post", data)
}

func TestTemplateCaching(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscape(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestRunTemplateCaching(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateCaching)
	mux.HandleFunc("/escape", TemplateAutoEscape)
	mux.HandleFunc("/nonescape", TemplateAutoEscapeDisabled)


	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}