package pkg

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path        string
	HandlerFunc gin.HandlerFunc
	Method      string
}

type Routes []Route
