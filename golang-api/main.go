package main

import (
	"golang_api/config"
	"golang_api/controller"
	"golang_api/helper"
	"golang_api/repository"
	"golang_api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

const CATEGORY_API_BASE_URL string ="/api/category"
const CATEGORY_API_PARAMS_ID string = "/api/category/:categoryId"

func main() {
	
	db := config.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.POST(CATEGORY_API_BASE_URL, categoryController.Create)
	router.GET(CATEGORY_API_BASE_URL, categoryController.FindAll)
	router.GET(CATEGORY_API_PARAMS_ID, categoryController.FindById)
	router.PUT(CATEGORY_API_PARAMS_ID, categoryController.Update)
	router.DELETE(CATEGORY_API_PARAMS_ID, categoryController.Delete)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}