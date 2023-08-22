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

const simpleHtml string = "simple.gohtml"

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil{
		panic(err)
	}
	t.ExecuteTemplate(w, simpleHtml, `<div className="container my-5">Hello Golang Web Templates</div>`)
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w,simpleHtml, "Hello HTML Template")
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

type PageContent struct{
	Title string
	Name string
}

type Hobby struct{
	Title string
	Hobbies []string
}

type Address struct{
	Street, City string
}

type People struct{
	Title string
	Nama string
	Alamat Address
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/layout.gohtml"))
	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title" : "Golang Template Data",
		"Content" : "Golang Web Template Data Example",
	})
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
		))
	t.ExecuteTemplate(w, "layout.gohtml", Page{
		Title: "Golang Template Data",
		Content: "Golang Web Template Data Struct",
	})
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", PageContent{
		Title: "Golang Template Action",
		Name: r.URL.Query().Get("name"),
	})
}

func TemplateActionComparator(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", struct{
		Title string
		FinalValue int
	}{
		Title: "Golang Template Action Comparator",
		FinalValue: 70,
	})
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", Hobby{
		Title: "Golang Template Range",
		Hobbies: []string{
			"Gaming", "Watching", "Coding",
		},
	})
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", People{
		Title: "Golang Template With",
		Nama: "Ichigo",
		Alamat: Address{
			Street: "District 38",
			City: "Rukongai",
		},
	})
}

func TestTemplateData(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateDataStruct)
	mux.HandleFunc("/range", TemplateActionRange)
	mux.HandleFunc("/with", TemplateActionWith)
	mux.HandleFunc("/map", TemplateDataMap)
	mux.HandleFunc("/action-if", TemplateActionIf)
	mux.HandleFunc("/action-comparator", TemplateActionComparator)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}