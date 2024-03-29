package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Ichigo" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Kurosaki",
	})
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Belajar Go-Lang",
	})
}

func TemplateFunctionMap(w http.ResponseWriter, r *http.Request){
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string{
			return strings.ToUpper(value)
		},
		"split": func(value string) []string{
			return strings.Split(value, " ")
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name | split}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dummy Name",
	})
}

func TestTemplateFunction(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateFunctionGlobal(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateFunctionCustomGlobal(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestRunTemplate(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateFunction)
	mux.HandleFunc("/global", TemplateFunctionGlobal)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}	

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}