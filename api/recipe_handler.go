package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeHandler struct {
	Service *services.RecipeService
}

var RecipeHandlerInstance *RecipeHandler

func (h *RecipeHandler) GetAllRecipes(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetAllRecipes()
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeHandler) GetRecipe(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetRecipe(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeHandler) CreateRecipe(ctx *gin.Context) {
	statusCode, body, err := h.Service.CreateRecipe(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeHandler) UpdateRecipe(ctx *gin.Context) {
	statusCode, body, err := h.Service.UpdateRecipe(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeHandler) DeleteRecipe(ctx *gin.Context) {
	statusCode, body, err := h.Service.DeleteRecipe(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *RecipeHandler) GetRecipeRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "recipes",
			HandlerFunc: h.GetAllRecipes,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "recipes/find-one/:id",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "recipes",
			HandlerFunc: h.CreateRecipe,
			Method:      "POST",
		},
		pkg.Route{
			Path:        "recipes",
			HandlerFunc: h.UpdateRecipe,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "recipes/:id",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
	}
}
