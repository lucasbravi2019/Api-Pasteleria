package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type IngredientPackageHandler struct {
	Service services.IngredientPackageService
}

type IngredientPackageHandlerInterface interface {
	AddPackageToIngredient(ctx *gin.Context)
	RemovePackageFromIngredients(ctx *gin.Context)
	GetIngredientPackageRoutes() []pkg.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) AddPackageToIngredient(ctx *gin.Context) {
	statusCode, err := h.Service.AddPackageToIngredient(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, err)
}

func (h *IngredientPackageHandler) RemovePackageFromIngredients(ctx *gin.Context) {
	statusCode, body, err := h.Service.RemovePackageFromIngredients(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, body, err)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/packages/:packageId/ingredients/:ingredientId",
			HandlerFunc: h.AddPackageToIngredient,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/packages/:packageId/ingredients",
			HandlerFunc: h.RemovePackageFromIngredients,
			Method:      "DELETE",
		},
	}
}
