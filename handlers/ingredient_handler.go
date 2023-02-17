package handlers

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type IngredientHandler struct {
	Service services.IngredientServiceInterface
}

type IngredientHandlerInterface interface {
	GetAllIngredients(w http.ResponseWriter, r *http.Request)
	CreateIngredient(w http.ResponseWriter, r *http.Request)
	UpdateIngredient(w http.ResponseWriter, r *http.Request)
	DeleteIngredient(w http.ResponseWriter, r *http.Request)
	ChangeIngredientPrice(w http.ResponseWriter, r *http.Request)
	GetIngredientRoutes() core.Routes
}

var IngredientHandlerInstance *IngredientHandler

func (h *IngredientHandler) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.GetAllIngredients()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *IngredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.CreateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *IngredientHandler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.UpdateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *IngredientHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.DeleteIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *IngredientHandler) ChangeIngredientPrice(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.ChangeIngredientPrice(r)
	core.EncodeJsonResponse(w, statusCode, body)
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
			Path:        "/ingredients/{id}",
			HandlerFunc: h.UpdateIngredient,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/ingredients/{id}/price",
			HandlerFunc: h.ChangeIngredientPrice,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/ingredients/{id}",
			HandlerFunc: h.DeleteIngredient,
			Method:      "DELETE",
		},
	}
}
