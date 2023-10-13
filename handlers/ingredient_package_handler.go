package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type IngredientPackageHandler struct {
	Service services.IngredientPackageService
}

type IngredientPackageHandlerInterface interface {
	AddPackageToIngredient(ctx *gin.Context)
	RemovePackageFromIngredients(ctx *gin.Context)
	GetIngredientPackageRoutes() []core.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) AddPackageToIngredient(ctx *gin.Context) {
	statusCode := h.Service.AddPackageToIngredient(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientPackageHandler) RemovePackageFromIngredients(ctx *gin.Context) {
	statusCode, body := h.Service.RemovePackageFromIngredients(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/packages/{packageId}/ingredients/{ingredientId}",
			HandlerFunc: h.AddPackageToIngredient,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/packages/{id}/ingredients",
			HandlerFunc: h.RemovePackageFromIngredients,
			Method:      "DELETE",
		},
	}
}
