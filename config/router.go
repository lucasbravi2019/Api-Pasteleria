package config

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/factory"
	"github.com/lucasbravi2019/pasteleria/middleware"
)

var apiRouterInstance *gin.Engine

func initRouter() *gin.Engine {
	if apiRouterInstance == nil {
		apiRouterInstance = gin.Default()
	}
	return apiRouterInstance
}

func registerRoutes(routes core.Routes) {
	for _, route := range routes {
		apiRouterInstance.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}

func corsConfig() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	apiRouterInstance.Use(cors.New(config))
}

func registerMiddleware(middleware gin.HandlerFunc) {
	apiRouterInstance.Use(middleware)
}

func StartApi() {
	initRouter()

	corsConfig()

	registerMiddleware(middleware.DatabaseCheckMiddleware())
	registerMiddleware(middleware.RequestLoggerMiddleware())

	registerRoutes(factory.GetRecipeHandlerInstance().GetRecipeRoutes())
	registerRoutes(factory.GetIngredientHandlerInstance().GetIngredientRoutes())
	registerRoutes(factory.GetPackageHandlerInstance().GetPackageRoutes())
	registerRoutes(factory.GetIngredientPackageHandlerInstance().GetIngredientPackageRoutes())
	registerRoutes(factory.GetRecipeIngredientHandlerInstance().GetRecipeIngredientRoutes())

	log.Fatal(apiRouterInstance.Run(":8080"))
}
