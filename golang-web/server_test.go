package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "OK!")
	})
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/hello/go", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Golang Web Hello World")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}