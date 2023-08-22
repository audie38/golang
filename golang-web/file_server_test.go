package golang_web

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T){
	directory := http.Dir("public")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	// mux.Handle("/static", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}