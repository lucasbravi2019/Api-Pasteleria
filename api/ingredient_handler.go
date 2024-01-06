package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type IngredientHandler struct {
	Service *services.IngredientService
}

var IngredientHandlerInstance *IngredientHandler

func (h *IngredientHandler) GetAllIngredients(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetAllIngredients()
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientHandler) CreateIngredient(ctx *gin.Context) {
	statusCode, body, err := h.Service.CreateIngredient(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientHandler) UpdateIngredient(ctx *gin.Context) {
	statusCode, body, err := h.Service.UpdateIngredient(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientHandler) DeleteIngredient(ctx *gin.Context) {
	statusCode, body, err := h.Service.DeleteIngredient(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *IngredientHandler) GetIngredientRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/ingredients",
			HandlerFunc: h.GetAllIngredients,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "/ingredients",
			HandlerFunc: h.CreateIngredient,
			Method:      "POST",
		},
		pkg.Route{
			Path:        "/ingredients",
			HandlerFunc: h.UpdateIngredient,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/ingredients/:id",
			HandlerFunc: h.DeleteIngredient,
			Method:      "DELETE",
		},
	}
}
