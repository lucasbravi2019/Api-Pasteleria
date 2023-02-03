package ingredients

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
)

type ingredientHandler struct {
	ingredientService IngredientService
}

type IngredientHandler interface {
	GetAllIngredients(w http.ResponseWriter, r *http.Request)
	CreateIngredient(w http.ResponseWriter, r *http.Request)
	UpdateIngredient(w http.ResponseWriter, r *http.Request)
	DeleteIngredient(w http.ResponseWriter, r *http.Request)
	GetIngredientRoutes() core.Routes
}

var ingredientHandlerInstance *ingredientHandler

func (h *ingredientHandler) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.ingredientService.GetAllIngredients()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *ingredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.ingredientService.CreateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *ingredientHandler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.ingredientService.UpdateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *ingredientHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.ingredientService.DeleteIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *ingredientHandler) GetIngredientRoutes() core.Routes {
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
			Path:        "/ingredients/{id}",
			HandlerFunc: h.DeleteIngredient,
			Method:      "DELETE",
		},
	}
}
