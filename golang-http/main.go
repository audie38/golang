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
		fmt.Fprint(w, "Hello World")
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