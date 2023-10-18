package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

func DatabaseCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.CheckDatabaseHealth()
		if pkg.HasError(err) {
			log.Println(err)
		}
	}
}
