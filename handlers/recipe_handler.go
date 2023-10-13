package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
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
	GetRecipeRoutes() core.Routes
}

var RecipeHandlerInstance *RecipeHandler

func (h *RecipeHandler) GetAllRecipes(ctx *gin.Context) {
	statusCode, body := h.Service.GetAllRecipes()
	core.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *RecipeHandler) GetRecipe(ctx *gin.Context) {
	statusCode, body := h.Service.GetRecipe(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *RecipeHandler) CreateRecipe(ctx *gin.Context) {
	err := h.Service.CreateRecipe(ctx.Request)
	statusCode := http.StatusCreated
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, err)
}

func (h *RecipeHandler) UpdateRecipeName(ctx *gin.Context) {
	statusCode := h.Service.UpdateRecipeName(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *RecipeHandler) DeleteRecipe(ctx *gin.Context) {
	statusCode := h.Service.DeleteRecipe(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *RecipeHandler) GetRecipeRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "recipes",
			HandlerFunc: h.GetAllRecipes,
			Method:      "GET",
		},
		core.Route{
			Path:        "recipes",
			HandlerFunc: h.CreateRecipe,
			Method:      "POST",
		},
		core.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.UpdateRecipeName,
			Method:      "PUT",
		},
		core.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		core.Route{
			Path:        "recipes/{id}",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
	}
}
