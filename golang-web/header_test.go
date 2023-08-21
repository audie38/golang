package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request){
	contentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer, contentType)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request){
	writer.Header().Add("x-powered-by", "Soul Society")
	fmt.Fprint(writer, "OK")
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	poweredBy := recorder.Header().Get("x-powered-by")
	fmt.Println(poweredBy)
}