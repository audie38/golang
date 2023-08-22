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

	t.ExecuteTemplate(w, "SIMPLE", `<div class="container my-5">Hello Golang</div>`)	
}

func TestTemplate(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}