package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
)

func DatabaseCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := core.CheckDatabaseHealth()
		if err != nil {
			log.Println(err)
		}
	}
}
