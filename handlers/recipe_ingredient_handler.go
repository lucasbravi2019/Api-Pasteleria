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
	AddIngredientToRecipe(ctx *gin.Context)
	GetRecipeIngredientRoutes() core.Routes
}

var RecipeIngredientHandlerInstance *RecipeIngredientHandler

func (h *RecipeIngredientHandler) AddIngredientToRecipe(ctx *gin.Context) {
	statusCode := h.Service.AddIngredientToRecipe(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
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
