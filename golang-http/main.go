package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const BASE_URL="localhost:8000"

//go:embed public
var public embed.FS

func main() {
	router := httprouter.New()
	directory, _ := fs.Sub(public, "public")
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}){
		fmt.Fprintf(w, "Panic %s", i)
	}

	router.ServeFiles("/files/*filepath", http.FS(directory))

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		queryParams := r.URL.Query().Get("name")
		if queryParams != ""{
			fmt.Fprintf(w, "Hello %s \n", queryParams)
		}
		fmt.Fprint(w, "Hello World")
	})
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprintf(w, "Product %s", params.ByName("id"))
	})
	router.GET("/product/:id/:slug", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprintf(w, "Product %s \n", params.ByName("id"))
		fmt.Fprintf(w, "Product Detail : %s", params.ByName("slug"))
	})
	router.GET("/image/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprintf(w, "Image: %s", params.ByName("image"))
	})
	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		panic("500 Internal Sever Error")
	})

	server := http.Server{
		Addr: BASE_URL,
		Handler: router,
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}