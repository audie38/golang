package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request){
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request){
	cookie, err := request.Cookie("token")
	if err != nil{
		fmt.Fprint(writer, "No Cookie")
	}else{
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetCookie)
	mux.HandleFunc("/cookie", SetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}

func TestSetCookie(t *testing.T){
	url := BASE_URL + "?name=Ichigo"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies{
		fmt.Printf("Cookie %s: %s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, BASE_URL, nil)
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "test"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	
}