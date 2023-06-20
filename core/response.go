package core

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

func EncodeJsonResponse(c *gin.Context, statusCode int, body interface{}, err error) {
	var response *Response = &Response{}

	if err != nil {
		response.Error = err.Error()
		c.JSON(statusCode, response)
		return
	}

	response.Body = body

	c.JSON(statusCode, response)
}
