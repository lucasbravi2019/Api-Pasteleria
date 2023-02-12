package recipes

import (
	"net/http"

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
	UpdateRecipeName(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)

	GetRecipeRoutes() core.Routes
}

func (h *recipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.GetAllRecipes()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.GetRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.CreateRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *recipeHandler) UpdateRecipeName(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.UpdateRecipeName(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *recipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.DeleteRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
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
			HandlerFunc: h.UpdateRecipeName,
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
