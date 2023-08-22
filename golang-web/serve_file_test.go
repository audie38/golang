package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		http.ServeFile(w, r, "public/index.html")
	}else{
		http.ServeFile(w, r, "public/404.html")
	}
}

//go:embed public/index.html
var resourceOk string

//go:embed public/404.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w, resourceOk)
	}else{
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFile(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}

func TestServeFileEmbed(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}