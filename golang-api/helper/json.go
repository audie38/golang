package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}