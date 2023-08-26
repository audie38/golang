package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const BASE_URL string = "http://localhost:8080"

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

func MultipleValueQueryParameter(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintln(writer, strings.Join(names, " "))
}

func TestQueryParameter(t *testing.T){
	url := BASE_URL + "?name=Ichigo"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T){
	url := BASE_URL + "?first_name=Ichgio&last_name=Kurosaki"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleValueQueryParameter(t *testing.T){
	url := BASE_URL + "?name=Shigekuni&name=Genryusai&name=Yamamoto"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}