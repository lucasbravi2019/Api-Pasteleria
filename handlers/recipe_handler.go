package handlers

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type RecipeHandler struct {
	Service services.RecipeService
}

type RecipeHandlerInterface interface {
	GetAllRecipes(w http.ResponseWriter, r *http.Request)
	GetRecipe(w http.ResponseWriter, r *http.Request)
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	UpdateRecipeName(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)
	GetRecipeRoutes() core.Routes
}

var RecipeHandlerInstance *RecipeHandler

func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.GetAllRecipes()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.GetRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.CreateRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeHandler) UpdateRecipeName(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.UpdateRecipeName(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.DeleteRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeHandler) GetRecipeRoutes() core.Routes {
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
