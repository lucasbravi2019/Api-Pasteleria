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
	decodeBody(r, &ingredient)

	core.EncodeJsonResponse(w, http.StatusCreated, h.ingredientService.CreateIngredient(ingredient))
}

func (h *ingredientHandler) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient Ingredient
	decodeBody(r, &ingredient)

	core.EncodeJsonResponse(w, http.StatusOK, h.ingredientService.UpdateIngredient(core.ConvertHexToObjectId(mux.Vars(r)["id"]), ingredient))
}

func (h *ingredientHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	core.EncodeJsonResponse(w, http.StatusOK, h.ingredientService.DeleteIngredient(core.ConvertHexToObjectId(mux.Vars(r)["id"])))
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

func decodeBody(r *http.Request, storeVar interface{}) {
	err := json.NewDecoder(r.Body).Decode(&storeVar)
	if err != nil {
		log.Fatal(err.Error())
	}
}
