package dto

type RecipeDTO struct {
	Id          int64                 `json:"id,omitempty"`
	Name        string                `json:"name"`
	Ingredients []RecipeIngredientDTO `json:"ingredients"`
	Price       float64               `json:"price"`
}

type RecipeCreationDTO struct {
	Name        string                   `json:"name"`
	Ingredients []RecipeIngredientIdsDTO `json:"ingredients"`
}

type RecipeUpdateDTO struct {
	Id          int64                    `json:"id"`
	Name        string                   `json:"name"`
	Ingredients []RecipeIngredientIdsDTO `json:"ingredients"`
}

type RecipeIngredientIdsDTO struct {
	IngredientId int64   `json:"ingredientId"`
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
}

func NewRecipeDTO(recipeId int64, recipeName string, recipePrice float64, ingredients []RecipeIngredientDTO) *RecipeDTO {
	return &RecipeDTO{
		Id:          recipeId,
		Name:        recipeName,
		Price:       recipePrice,
		Ingredients: ingredients,
	}
}
