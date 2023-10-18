package config

import (
	"time"

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
	router := GetRouter()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}

func StartApi() {
	LoadEnv()
	db.QueryLoader()
	RegisterRoutes(factory.GetRecipeHandlerInstance().GetRecipeRoutes())
	RegisterRoutes(factory.GetIngredientHandlerInstance().GetIngredientRoutes())
	RegisterRoutes(factory.GetPackageHandlerInstance().GetPackageRoutes())
	RegisterRoutes(factory.GetIngredientPackageHandlerInstance().GetIngredientPackageRoutes())

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
