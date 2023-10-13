package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
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
	GetIngredientRoutes() pkg.Routes
}

var IngredientHandlerInstance *IngredientHandler

func (h *IngredientHandler) GetAllIngredients(ctx *gin.Context) {
	statusCode, body := h.Service.GetAllIngredients()
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *IngredientHandler) CreateIngredient(ctx *gin.Context) {
	statusCode := h.Service.CreateIngredient(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) UpdateIngredient(ctx *gin.Context) {
	statusCode := h.Service.UpdateIngredient(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) DeleteIngredient(ctx *gin.Context) {
	statusCode := h.Service.DeleteIngredient(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *IngredientHandler) ChangeIngredientPrice(ctx *gin.Context) {
	statusCode := h.Service.ChangeIngredientPrice(ctx.Request)
	pkg.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
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
			Path:        "/ingredients/:ingredientId",
			HandlerFunc: h.UpdateIngredient,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/ingredients/:ingredientId/price",
			HandlerFunc: h.ChangeIngredientPrice,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/ingredients/:ingredientId",
			HandlerFunc: h.DeleteIngredient,
			Method:      "DELETE",
		},
	}
}
