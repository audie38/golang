package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr: "localhost:8000",
		Handler: router,
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}