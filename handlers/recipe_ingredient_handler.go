package handlers

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type RecipeIngredientHandler struct {
	Service services.RecipeIngredientService
}

type RecipeIngredientHandlerInterface interface {
	AddIngredientToRecipe(w http.ResponseWriter, r *http.Request)
	RemoveIngredientFromRecipe(w http.ResponseWriter, r *http.Request)
	GetRecipeIngredientRoutes() core.Routes
}

var RecipeIngredientHandlerInstance *RecipeIngredientHandler

func (h *RecipeIngredientHandler) AddIngredientToRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode := h.Service.AddIngredientToRecipe(r)
	core.EncodeJsonResponse(w, statusCode, nil)
}

func (h *RecipeIngredientHandler) RemoveIngredientFromRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.RemoveIngredientFromRecipe(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *RecipeIngredientHandler) GetRecipeIngredientRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/ingredients/{ingredientId}/recipes/{recipeId}",
			HandlerFunc: h.AddIngredientToRecipe,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/recipes/remove-ingredient",
			HandlerFunc: h.RemoveIngredientFromRecipe,
			Method:      "PUT",
		},
	}
}
