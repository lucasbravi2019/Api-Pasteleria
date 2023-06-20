package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type RecipeIngredientHandler struct {
	Service services.RecipeIngredientService
}

type RecipeIngredientHandlerInterface interface {
	AddIngredientToRecipe(c *gin.Context)
	RemoveIngredientFromRecipe(c *gin.Context)
	GetRecipeIngredientRoutes() core.Routes
}

var RecipeIngredientHandlerInstance *RecipeIngredientHandler

func (h *RecipeIngredientHandler) AddIngredientToRecipe(c *gin.Context) {
	statusCode, err := h.Service.AddIngredientToRecipe(c)
	core.EncodeJsonResponse(c, statusCode, nil, err)
}

func (h *RecipeIngredientHandler) RemoveIngredientFromRecipe(c *gin.Context) {
	statusCode, body, err := h.Service.RemoveIngredientFromRecipe(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *RecipeIngredientHandler) GetRecipeIngredientRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/ingredients/:ingredientId/recipes/:recipeId",
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
