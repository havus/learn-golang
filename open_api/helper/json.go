package helper

import (
	"encoding/json"
	"net/http"
	"open_api/model/web"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err 		:= decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	errEncoder := encoder.Encode(response)
	PanicIfError(errEncoder)
}