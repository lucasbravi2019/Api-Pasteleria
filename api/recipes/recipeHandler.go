package recipes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
)

type recipeHandler struct {
	service RecipeService
}

var recipeHandlerInstance *recipeHandler

type RecipeHandler interface {
	GetAllRecipes(w http.ResponseWriter, r *http.Request)
	GetRecipe(w http.ResponseWriter, r *http.Request)
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	UpdateRecipe(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)
	GetRecipeRoutes() core.Routes
}

func (h *recipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := h.service.GetAllRecipes()

	core.EncodeJsonResponse(w, http.StatusCreated, recipes)
}

func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := h.service.GetRecipe(core.ConvertHexToObjectId(mux.Vars(r)["id"]))

	core.EncodeJsonResponse(w, http.StatusOK, recipe)
}

func (h *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h *recipeHandler) GetRecipeRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/recipes",
			HandlerFunc: h.GetAllRecipes,
			Method:      "GET",
		},
		core.Route{
			Path:        "/recipes",
			HandlerFunc: h.CreateRecipe,
			Method:      "POST",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.UpdateRecipe,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
	}
}
