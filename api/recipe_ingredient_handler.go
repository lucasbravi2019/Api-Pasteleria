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
	statusCode, body, err := h.Service.AddIngredientToRecipe(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeIngredientHandler) GetRecipeIngredientRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/ingredients/:id/recipes/:recipeId",
			HandlerFunc: h.AddIngredientToRecipe,
			Method:      "PUT",
		},
	}
}
