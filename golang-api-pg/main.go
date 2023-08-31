package main

import (
	"golang_api_pg/config"
	"golang_api_pg/controller"
	"golang_api_pg/helper"
	"golang_api_pg/middleware"
	"golang_api_pg/repository"
	"golang_api_pg/service"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	repo := repository.NewCategoryRepository()
	service := service.NewCategoryService(repo, db, validate)
	controller := controller.NewCategoryController(service)
	router := config.NewRouter(controller)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}