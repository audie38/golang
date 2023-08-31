package config

import (
	"golang_api_pg/controller"
	"golang_api_pg/exception"
	"golang_api_pg/helper"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router{
	router := httprouter.New()

	router.POST(helper.CATEGORY_API_BASE_URL, categoryController.Create)
	router.GET(helper.CATEGORY_API_BASE_URL, categoryController.FindAll)
	router.GET(helper.CATEGORY_API_PARAMS_ID, categoryController.FindById)
	router.PUT(helper.CATEGORY_API_PARAMS_ID, categoryController.Update)
	router.DELETE(helper.CATEGORY_API_PARAMS_ID, categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}