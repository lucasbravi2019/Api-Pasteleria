package handlers

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type IngredientPackageHandler struct {
	Service services.IngredientPackageService
}

type IngredientPackageHandlerInterface interface {
	AddPackageToIngredient(w http.ResponseWriter, r *http.Request)
	RemovePackageFromIngredients(w http.ResponseWriter, r *http.Request)
	GetIngredientPackageRoutes() []core.Route
}

var IngredientPackageHandlerInstance *IngredientPackageHandler

func (h *IngredientPackageHandler) AddPackageToIngredient(w http.ResponseWriter, r *http.Request) {
	statusCode := h.Service.AddPackageToIngredient(r)
	core.EncodeJsonResponse(w, statusCode, nil)
}

func (h *IngredientPackageHandler) RemovePackageFromIngredients(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.RemovePackageFromIngredients(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *IngredientPackageHandler) GetIngredientPackageRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/packages/{packageId}/ingredients/{ingredientId}",
			HandlerFunc: h.AddPackageToIngredient,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/packages/{id}/ingredients",
			HandlerFunc: h.RemovePackageFromIngredients,
			Method:      "DELETE",
		},
	}
}
