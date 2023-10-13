package core

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

func EncodeJsonResponse(w gin.ResponseWriter, statusCode int, body interface{}, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	var response *Response = &Response{}

	if statusCode == http.StatusInternalServerError ||
		statusCode == http.StatusBadRequest ||
		statusCode == http.StatusNotFound {
		response.Error = "Ocurrio un error al realizar la operacion"
	}

	if statusCode == http.StatusOK && body == nil {
		body = "OK"
	}

	if body != nil {
		response.Body = body
		json.NewEncoder(w).Encode(response)
	}
}
