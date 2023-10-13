package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeIngredientHandler struct {
	Service services.RecipeIngredientService
}

type RecipeIngredientHandlerInterface interface {
	AddIngredientToRecipe(ctx *gin.Context)
	GetRecipeIngredientRoutes() pkg.Routes
}

var RecipeIngredientHandlerInstance *RecipeIngredientHandler

func (h *RecipeIngredientHandler) AddIngredientToRecipe(ctx *gin.Context) {
	statusCode, err := h.Service.AddIngredientToRecipe(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, err)
}

func (h *RecipeIngredientHandler) GetRecipeIngredientRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/ingredients/:ingredientId/recipes/:recipeId",
			HandlerFunc: h.AddIngredientToRecipe,
			Method:      "PUT",
		},
	}
}
