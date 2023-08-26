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
//go:embed public/404.html
var notFound string

type LogMiddleware struct{
	Handler http.Handler
}

type ErrorMiddleware struct{
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Start Log")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Finish Log")
}

func (middleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	defer func (){
		message := recover()
		if message != nil{
			fmt.Printf("Recovered Error: %s \n", message)
		}
	}()
	middleware.Handler.ServeHTTP(writer, request)
}

func main() {
	router := httprouter.New()
	directory, _ := fs.Sub(public, "public")
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}){
		fmt.Fprintf(w, "Panic %s", i)
	}
	router.NotFound = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, notFound)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Method Not Allowed")
	})
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

	router.POST("/testing", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprint(w, "Post Method")
	})


	logMD := &LogMiddleware{router}
	errMD := &ErrorMiddleware{logMD}

	server := http.Server{
		Addr: BASE_URL,
		Handler: errMD,
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}