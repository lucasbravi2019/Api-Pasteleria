package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonRequest *json.RawMessage = &json.RawMessage{}

		err := json.NewDecoder(c.Request.Body).Decode(jsonRequest)
		if pkg.HasError(err) && err.Error() != "EOF" {
			log.Println(err.Error())
		}
		c.Request.Body.Close()

		body, err := jsonRequest.MarshalJSON()
		if pkg.HasError(err) {
			log.Println(err)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		log.Printf("URL Request: %s: %s%s\n", c.Request.Method, c.Request.Host, c.Request.URL.Path)
		stringBody := string(body)
		if stringBody != "" {
			log.Println(stringBody)
		}
	}
}
