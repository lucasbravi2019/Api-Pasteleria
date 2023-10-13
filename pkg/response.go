package pkg

import (
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

	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(statusCode)
		return
	}

	response.Body = body
}
