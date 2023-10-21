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
	FindAllIngredientPackages(ctx *gin.Context)
	UpdateIngredientPackages(ctx *gin.Context)
	GetIngredientPackageRoutes() []pkg.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) FindAllIngredientPackages(ctx *gin.Context) {
	statusCode, body, err := h.Service.FindAllIngredientPackages(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientPackageHandler) UpdateIngredientPackages(ctx *gin.Context) {
	statusCode, body, err := h.Service.UpdateIngredientPackages(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/ingredients/packages",
			HandlerFunc: h.FindAllIngredientPackages,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "/ingredients/packages",
			HandlerFunc: h.UpdateIngredientPackages,
			Method:      "PUT",
		},
	}
}
