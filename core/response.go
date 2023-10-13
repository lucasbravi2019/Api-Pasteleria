package core

import (
<<<<<<< HEAD
	"encoding/json"
	"net/http"

=======
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

<<<<<<< HEAD
func EncodeJsonResponse(w gin.ResponseWriter, statusCode int, body interface{}, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
=======
func EncodeJsonResponse(c *gin.Context, statusCode int, body interface{}, err error) {
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	var response *Response = &Response{}

	if err != nil {
		response.Error = err.Error()
		c.JSON(statusCode, response)
		return
	}

	response.Body = body

	c.JSON(statusCode, response)
}
