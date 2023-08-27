package exception

import (
	"golang_api/helper"
	"golang_api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

const CONTENT_TYPE string = "Content-Type"
const APP_JSON string = "application/json"

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}){
	if notFoundError(w, r, err){
		return 
	}

	if validationErrors(w, r, err){
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}){
	w.Header().Add(CONTENT_TYPE, APP_JSON)
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data: err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, ok := err.(NotFoundError)
	if ok{
		w.Header().Add(CONTENT_TYPE, APP_JSON)
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Data: exception.Error,
		}
	
		helper.WriteToResponseBody(w, webResponse)
		return true
	}else{
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, ok := err.(validator.ValidationErrors)
	if ok{
		w.Header().Add(CONTENT_TYPE, APP_JSON)
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Data: exception.Error(),
		}
	
		helper.WriteToResponseBody(w, webResponse)
		return true
	}else{
		return false
	}
}