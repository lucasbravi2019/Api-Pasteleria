package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/api/middleware"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/factory"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

var apiRouterInstance *gin.Engine

func GetRouter() *gin.Engine {
	if apiRouterInstance == nil {
		apiRouterInstance = gin.Default()
	}
	return apiRouterInstance
}

func RegisterRoutes(routes pkg.Routes) {
	r := GetRouter()
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}

func StartApi() {
	r := GetRouter()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLoggerMiddleware())
	r.Use(middleware.DatabaseCheckMiddleware())

	LoadEnv()
	db.QueryLoader()
	RegisterRoutes(factory.GetRecipeHandlerInstance().GetRecipeRoutes())
	RegisterRoutes(factory.GetIngredientHandlerInstance().GetIngredientRoutes())
	RegisterRoutes(factory.GetPackageHandlerInstance().GetPackageRoutes())
	RegisterRoutes(factory.GetIngredientPackageHandlerInstance().GetIngredientPackageRoutes())
	RegisterRoutes(factory.GetRecipeIngredientHandlerInstance().GetAllRecipeIngredientRoutes())

	r.Run(":8080")
}
