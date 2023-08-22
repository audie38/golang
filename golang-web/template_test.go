package golang_web

import (
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

func TestTemplateLayout(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", SimpleHTML)
	mux.HandleFunc("/template", SimpleHTMLFile)

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