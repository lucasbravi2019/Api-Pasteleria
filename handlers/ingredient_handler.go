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
	GetAllIngredients(c *gin.Context)
	CreateIngredient(c *gin.Context)
	UpdateIngredient(c *gin.Context)
	DeleteIngredient(c *gin.Context)
	ChangeIngredientPrice(c *gin.Context)
	GetIngredientRoutes() core.Routes
}

var IngredientHandlerInstance *IngredientHandler

func (h *IngredientHandler) GetAllIngredients(c *gin.Context) {
	statusCode, body, err := h.Service.GetAllIngredients()
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *IngredientHandler) CreateIngredient(c *gin.Context) {
	statusCode, body, err := h.Service.CreateIngredient(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *IngredientHandler) UpdateIngredient(c *gin.Context) {
	statusCode, body, err := h.Service.UpdateIngredient(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *IngredientHandler) DeleteIngredient(c *gin.Context) {
	statusCode, body, err := h.Service.DeleteIngredient(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *IngredientHandler) ChangeIngredientPrice(c *gin.Context) {
	statusCode, body, err := h.Service.ChangeIngredientPrice(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
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
