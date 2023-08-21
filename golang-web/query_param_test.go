package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(writer, "Hello")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request){
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}


func TestQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Ichigo", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?first_name=Ichgio&last_name=Kurosaki", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}