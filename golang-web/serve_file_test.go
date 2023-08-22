package golang_web

import (
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