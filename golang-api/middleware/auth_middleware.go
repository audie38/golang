package middleware

import (
	"golang_api/helper"
	"golang_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	if "SECRET" == request.Header.Get("X-API-KEY"){
		middleware.Handler.ServeHTTP(writer, request)
	}else{
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UnAuthorized",
		}
	
		helper.WriteToResponseBody(writer, webResponse)
	}
}