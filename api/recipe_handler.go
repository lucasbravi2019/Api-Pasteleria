package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeHandler struct {
	Service services.RecipeService
}

type RecipeHandlerInterface interface {
	GetAllRecipes(ctx *gin.Context)
	GetRecipe(ctx *gin.Context)
	CreateRecipe(ctx *gin.Context)
	UpdateRecipeName(ctx *gin.Context)
	DeleteRecipe(ctx *gin.Context)
	GetRecipeRoutes() pkg.Routes
}

var RecipeHandlerInstance *RecipeHandler

func (h *RecipeHandler) GetAllRecipes(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetAllRecipes()
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, body, err)
}

func (h *RecipeHandler) GetRecipe(ctx *gin.Context) {
	statusCode, body := h.Service.GetRecipe(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *RecipeHandler) CreateRecipe(ctx *gin.Context) {
	err := h.Service.CreateRecipe(ctx.Request)
	statusCode := http.StatusCreated
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, err)
}

func (h *RecipeHandler) UpdateRecipeName(ctx *gin.Context) {
	statusCode := h.Service.UpdateRecipeName(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *RecipeHandler) DeleteRecipe(ctx *gin.Context) {
	statusCode := h.Service.DeleteRecipe(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *RecipeHandler) GetRecipeRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "recipes",
			HandlerFunc: h.GetAllRecipes,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "recipes",
			HandlerFunc: h.CreateRecipe,
			Method:      "POST",
		},
		pkg.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.UpdateRecipeName,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
	}
}
