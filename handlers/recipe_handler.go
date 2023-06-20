package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type RecipeHandler struct {
	Service services.RecipeService
}

type RecipeHandlerInterface interface {
	GetAllRecipes(c *gin.Context)
	GetRecipe(c *gin.Context)
	CreateRecipe(c *gin.Context)
	UpdateRecipeName(c *gin.Context)
	DeleteRecipe(c *gin.Context)
	GetRecipeRoutes() core.Routes
}

var RecipeHandlerInstance *RecipeHandler

func (h *RecipeHandler) GetAllRecipes(c *gin.Context) {
	statusCode, body, err := h.Service.GetAllRecipes()
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *RecipeHandler) GetRecipe(c *gin.Context) {
	statusCode, body, err := h.Service.GetRecipe(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
	statusCode, body, err := h.Service.CreateRecipe(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *RecipeHandler) UpdateRecipeName(c *gin.Context) {
	statusCode, body, err := h.Service.UpdateRecipeName(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *RecipeHandler) DeleteRecipe(c *gin.Context) {
	statusCode, body, err := h.Service.DeleteRecipe(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
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
			Path:        "/recipes/update-name",
			HandlerFunc: h.UpdateRecipeName,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/recipes/find-one",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		core.Route{
			Path:        "/recipes/delete-one/:id",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
	}
}
