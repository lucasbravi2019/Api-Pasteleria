package ingredients

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	core.EncodeJsonResponse(w, http.StatusOK, h.ingredientService.GetAllIngredients())
}

func (h *ingredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient Ingredient
	errExists := decodeBody(r, &ingredient)

	if errExists {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}
	oid, err := h.ingredientService.CreateIngredient(ingredient)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusCreated, oid)
	}
}

func (h *ingredientHandler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient Ingredient
	errExists := decodeBody(r, &ingredient)

	if errExists {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	oid, err := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	id, err := h.ingredientService.UpdateIngredient(oid, ingredient)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusOK, id)
	}

}

func (h *ingredientHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	oid, err := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	err = h.ingredientService.DeleteIngredient(oid)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusNoContent, nil)
	}
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

func decodeBody(r *http.Request, storeVar interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(&storeVar)
	if err != nil {
		log.Println(err.Error())
		return true
	}

	if core.Validate(storeVar) {
		return true
	}

	return false
}
