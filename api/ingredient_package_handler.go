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
	FindAllIngredientPackages(ctx *gin.Context)
	GetIngredientPackageRoutes() []pkg.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) AddPackageToIngredient(ctx *gin.Context) {
	statusCode, body, err := h.Service.AddPackageToIngredient(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientPackageHandler) RemovePackageFromIngredients(ctx *gin.Context) {
	statusCode, body, err := h.Service.RemovePackageFromIngredients(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientPackageHandler) FindAllIngredientPackages(ctx *gin.Context) {
	statusCode, body, err := h.Service.FindAllIngredientPackages(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/packages/:id/ingredients/:ingredientId",
			HandlerFunc: h.AddPackageToIngredient,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/packages/:id/ingredients",
			HandlerFunc: h.RemovePackageFromIngredients,
			Method:      "DELETE",
		},
		pkg.Route{
			Path:        "/ingredients/:id/packages",
			HandlerFunc: h.FindAllIngredientPackages,
			Method:      "GET",
		},
	}
}
