package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil{
		panic(err)
	}

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostFormValue("lastName") // alternative

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("firstName=Ichigo&lastName=Kurosaki")
	request := httptest.NewRequest(http.MethodPost, BASE_URL, requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}