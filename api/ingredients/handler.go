package ingredients

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
)

type handler struct {
	service IngredientService
}

type IngredientHandler interface {
	GetAllIngredients(w http.ResponseWriter, r *http.Request)
	CreateIngredient(w http.ResponseWriter, r *http.Request)
	UpdateIngredient(w http.ResponseWriter, r *http.Request)
	DeleteIngredient(w http.ResponseWriter, r *http.Request)
	AddIngredientToRecipe(w http.ResponseWriter, r *http.Request)
	AddPackageToIngredient(w http.ResponseWriter, r *http.Request)
	RemovePackageFromIngredients(w http.ResponseWriter, r *http.Request)
	ChangeIngredientPrice(w http.ResponseWriter, r *http.Request)
	GetIngredientRoutes() core.Routes
}

var ingredientHandlerInstance *handler

func (h *handler) GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.GetAllIngredients()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.CreateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.UpdateIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.DeleteIngredient(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) AddIngredientToRecipe(w http.ResponseWriter, r *http.Request) {
	statusCode := h.service.AddIngredientToRecipe(r)
	core.EncodeJsonResponse(w, statusCode, nil)
}

func (h *handler) ChangeIngredientPrice(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.ChangeIngredientPrice(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) GetIngredientRoutes() core.Routes {
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
		core.Route{
			Path:        "/ingredients/{ingredientId}/recipes/{recipeId}",
			HandlerFunc: h.AddIngredientToRecipe,
			Method:      "PUT",
		},
	}
}
