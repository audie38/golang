package exception

import (
	"golang_api/helper"
	"golang_api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

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
	w.Header().Add(helper.CONTENT_TYPE, helper.APP_JSON)
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: helper.INTERNAL_SERVER_ERROR,
		Data: err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, ok := err.(NotFoundError)
	if ok{
		w.Header().Add(helper.CONTENT_TYPE, helper.APP_JSON)
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: helper.NOT_FOUND_ERROR,
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
		w.Header().Add(helper.CONTENT_TYPE, helper.APP_JSON)
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: helper.BAD_REQUEST_ERROR,
			Data: exception.Error(),
		}
	
		helper.WriteToResponseBody(w, webResponse)
		return true
	}else{
		return false
	}
}