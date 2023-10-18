package pkg

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error"`
	Body  interface{} `json:"body"`
}

func EncodeJsonResponse(c *gin.Context, statusCode int, body interface{}, err error) {
	response := &Response{}

	if HasError(err) {
		response.Error = err.Error()
		c.JSON(statusCode, response)
		return
	}

	if body != nil {
		response.Body = body
	}

	c.JSON(statusCode, response)
}
