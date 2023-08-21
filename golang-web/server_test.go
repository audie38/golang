package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}