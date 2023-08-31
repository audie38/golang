package middleware

import (
	"golang_api_pg/helper"
	"golang_api_pg/model/web"
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
		writer.Header().Add(helper.CONTENT_TYPE, helper.APP_JSON)
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: helper.UNAUTHORIZED_ERROR,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}