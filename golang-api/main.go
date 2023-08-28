package main

import (
	"golang_api/config"
	"golang_api/controller"
	"golang_api/helper"
	"golang_api/middleware"
	"golang_api/repository"
	"golang_api/service"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	
	db := config.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := config.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}