package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const BASE_URL="localhost:8000"

func main() {
	router := httprouter.New()
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

	server := http.Server{
		Addr: BASE_URL,
		Handler: router,
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}