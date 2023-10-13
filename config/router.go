package config

import (
<<<<<<< HEAD
	"time"
=======
	"log"
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/factory"
	"github.com/lucasbravi2019/pasteleria/middleware"
)

var apiRouterInstance *gin.Engine

<<<<<<< HEAD
func GetRouter() *gin.Engine {
=======
func initRouter() *gin.Engine {
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	if apiRouterInstance == nil {
		apiRouterInstance = gin.Default()
	}
	return apiRouterInstance
}

func registerRoutes(routes core.Routes) {
	for _, route := range routes {
<<<<<<< HEAD
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
=======
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
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
}
