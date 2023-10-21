package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeIngredientHandler struct {
	Service services.RecipeIngredientService
}

var RecipeIngredientHandlerInstance *RecipeIngredientHandler

type RecipeIngredientInterface interface {
	GetAllRecipeIngredients(ctx *gin.Context)
	UpdateRecipeIngredients(ctx *gin.Context)
	GetAllRecipeIngredientRoutes(ctx *gin.Context) pkg.Routes
}

func (h *RecipeIngredientHandler) GetAllRecipeIngredients(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetAllRecipeIngredients(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeIngredientHandler) UpdateRecipeIngredients(ctx *gin.Context) {
	statusCode, body, err := h.Service.UpdateRecipeIngredients(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeIngredientHandler) GetAllRecipeIngredientRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "recipes/ingredients",
			HandlerFunc: h.GetAllRecipeIngredients,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "recipes/ingredients",
			HandlerFunc: h.UpdateRecipeIngredients,
			Method:      "PUT",
		},
	}
}
