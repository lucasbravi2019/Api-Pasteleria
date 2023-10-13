package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/factory"
	"github.com/lucasbravi2019/pasteleria/middleware"
)

var apiRouterInstance *gin.Engine

func GetRouter() *gin.Engine {
	if apiRouterInstance == nil {
		apiRouterInstance = gin.Default()
	}
	return apiRouterInstance
}

func RegisterRoutes(routes core.Routes) {
	router := GetRouter()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}

func StartApi() {
	LoadEnv()
	core.QueryLoader()
	RegisterRoutes(factory.GetRecipeHandlerInstance().GetRecipeRoutes())
	RegisterRoutes(factory.GetIngredientHandlerInstance().GetIngredientRoutes())
	RegisterRoutes(factory.GetPackageHandlerInstance().GetPackageRoutes())
	RegisterRoutes(factory.GetIngredientPackageHandlerInstance().GetIngredientPackageRoutes())
	RegisterRoutes(factory.GetRecipeIngredientHandlerInstance().GetRecipeIngredientRoutes())

	r := GetRouter()
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:           1 * time.Hour,
	}))

	r.Use(middleware.RequestLoggerMiddleware())
	r.Use(middleware.DatabaseCheckMiddleware())
	r.Run(":8080")
}
