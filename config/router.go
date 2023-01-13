package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
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
		router.Path(route.Path).HandlerFunc(middleware.DatabaseCheckMiddleware(route.HandlerFunc)).Methods(route.Method)
	}
}

func StartApi() {
	RegisterRoutes(recipes.GetRecipeHandlerInstance().GetRecipeRoutes())
	RegisterRoutes(ingredients.GetIngredientHandlerInstance().GetIngredientRoutes())

	http.ListenAndServe("localhost:8080", GetRouter())
}
