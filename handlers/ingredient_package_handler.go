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
	AddPackageToIngredient(c *gin.Context)
	RemovePackageFromIngredients(c *gin.Context)
	GetIngredientPackageRoutes() []core.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) AddPackageToIngredient(c *gin.Context) {
	statusCode, err := h.Service.AddPackageToIngredient(c)
	core.EncodeJsonResponse(c, statusCode, nil, err)
}

func (h *IngredientPackageHandler) RemovePackageFromIngredients(c *gin.Context) {
	statusCode, body, err := h.Service.RemovePackageFromIngredients(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/packages/:packageId/ingredients/:ingredientId",
			HandlerFunc: h.AddPackageToIngredient,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/packages/:packageId/ingredients",
			HandlerFunc: h.RemovePackageFromIngredients,
			Method:      "DELETE",
		},
	}
}
