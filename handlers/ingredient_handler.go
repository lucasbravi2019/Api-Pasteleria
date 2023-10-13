package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type IngredientHandler struct {
	Service services.IngredientServiceInterface
}

type IngredientHandlerInterface interface {
	GetAllIngredients(ctx *gin.Context)
	CreateIngredient(ctx *gin.Context)
	UpdateIngredient(ctx *gin.Context)
	DeleteIngredient(ctx *gin.Context)
	ChangeIngredientPrice(ctx *gin.Context)
	GetIngredientRoutes() core.Routes
}

var IngredientHandlerInstance *IngredientHandler

func (h *IngredientHandler) GetAllIngredients(ctx *gin.Context) {
	statusCode, body := h.Service.GetAllIngredients()
	core.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *IngredientHandler) CreateIngredient(ctx *gin.Context) {
	statusCode := h.Service.CreateIngredient(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) UpdateIngredient(ctx *gin.Context) {
	statusCode := h.Service.UpdateIngredient(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) DeleteIngredient(ctx *gin.Context) {
	statusCode := h.Service.DeleteIngredient(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) ChangeIngredientPrice(ctx *gin.Context) {
	statusCode := h.Service.ChangeIngredientPrice(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) GetIngredientRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/ingredients",
			HandlerFunc: h.GetAllIngredients,
			Method:      "GET",
		},
		core.Route{
			Path:        "/ingredients",
			HandlerFunc: h.CreateIngredient,
			Method:      "POST",
		},
		core.Route{
			Path:        "/ingredients/:ingredientId",
			HandlerFunc: h.UpdateIngredient,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/ingredients/:ingredientId/price",
			HandlerFunc: h.ChangeIngredientPrice,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/ingredients/:ingredientId",
			HandlerFunc: h.DeleteIngredient,
			Method:      "DELETE",
		},
	}
}
