package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/recipes"
	"github.com/lucasbravi2019/pasteleria/core"
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
		router.Path(route.Path).HandlerFunc(route.HandlerFunc).Methods(route.Method)
	}
}

func StartApi() {
	RegisterRoutes(recipes.GetRecipeHandlerInstance().GetRecipeRoutes())

	http.ListenAndServe("localhost:8080", GetRouter())
}
