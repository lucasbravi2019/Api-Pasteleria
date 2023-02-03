package recipes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
)

type recipeHandler struct {
	service RecipeService
}

var recipeHandlerInstance *recipeHandler

type RecipeHandler interface {
	GetAllRecipes(w http.ResponseWriter, r *http.Request)
	GetRecipe(w http.ResponseWriter, r *http.Request)
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	UpdateRecipe(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)
	AddIngredientToRecipe(w http.ResponseWriter, r *http.Request)
	GetRecipeRoutes() core.Routes
}

func (h *recipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetAllRecipes()

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusOK, recipes)
	}
}

func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	oid, err := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	recipe, err := h.service.GetRecipe(oid)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusOK, recipe)
	}
}

func (h *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipeName RecipeName
	errExists := decodeBody(r, &recipeName)

	if errExists {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	oid, err := h.service.CreateRecipe(recipeName)

	if err != nil {
		log.Println(err)
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		log.Println(oid)
		core.EncodeJsonResponse(w, http.StatusCreated, oid)
	}
}

func (h *recipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	errExists := decodeBody(r, &recipe)

	if errExists {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	oid, err := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	id, err := h.service.UpdateRecipe(oid, recipe)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusOK, id)
	}
}

func (h *recipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	oid, err := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	err = h.service.DeleteRecipe(oid)

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusInternalServerError, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusNoContent, nil)
	}
}

func (h *recipeHandler) AddIngredientToRecipe(w http.ResponseWriter, r *http.Request) {
	var ingredientDetails IngredientDetails
	errExists := decodeBody(r, &ingredientDetails)

	if errExists {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	recipeId, err := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	ingredientId, err := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if err != nil {
		core.EncodeJsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	err, statusCode, recipe := h.service.AddIngredientToRecipe(recipeId, ingredientId, ingredientDetails)

	if err != nil {
		core.EncodeJsonResponse(w, statusCode, nil)
	} else {
		core.EncodeJsonResponse(w, http.StatusOK, recipe)
	}
}

func (h *recipeHandler) GetRecipeRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/recipes",
			HandlerFunc: h.GetAllRecipes,
			Method:      "GET",
		},
		core.Route{
			Path:        "/recipes",
			HandlerFunc: h.CreateRecipe,
			Method:      "POST",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.UpdateRecipe,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.GetRecipe,
			Method:      "GET",
		},
		core.Route{
			Path:        "/recipes/{id}",
			HandlerFunc: h.DeleteRecipe,
			Method:      "DELETE",
		},
		core.Route{
			Path:        "/recipes/{recipeId}/ingredient/{ingredientId}",
			HandlerFunc: h.AddIngredientToRecipe,
			Method:      "PUT",
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
