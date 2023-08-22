package golang_web

import (
	"embed"
	"io/fs"
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

//go:embed public
var public embed.FS

func TestFileServerGolangEmbed(t *testing.T){
	directory, _ := fs.Sub(public, "public")
	fileServer := http.FileServer(http.FS(directory))
	// fileServer := http.FileServer(http.FS(public))
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
	}
}