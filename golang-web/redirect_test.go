package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello Redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}