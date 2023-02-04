package config

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"github.com/lucasbravi2019/pasteleria/api/recipes"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/middleware"
)

var apiRouterInstance *mux.Router

func GetRouter() *mux.Router {
	if apiRouterInstance == nil {
		apiRouterInstance = mux.NewRouter()
	}
	return apiRouterInstance
}

func RegisterRoutes(routes core.Routes) {
	router := GetRouter()
	for _, route := range routes {
		router.
			Path(route.Path).
			HandlerFunc(
				middleware.RequestLoggerMiddleware(
					middleware.DatabaseCheckMiddleware(route.HandlerFunc))).
			Methods(route.Method)
	}
}

func StartApi() {
	RegisterRoutes(recipes.GetRecipeHandlerInstance().GetRecipeRoutes())
	RegisterRoutes(ingredients.GetIngredientHandlerInstance().GetIngredientRoutes())
	RegisterRoutes(packages.GetPackageHandlerInstance().GetPackageRoutes())

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(credentials, methods, ttl, origins)(GetRouter())))
}
