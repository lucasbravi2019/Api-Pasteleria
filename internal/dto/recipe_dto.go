package dto

type RecipeDTO struct {
	Id          int64                 `json:"id"`
	Name        string                `json:"name"`
	Ingredients []RecipeIngredientDTO `json:"ingredients"`
	Price       float64               `json:"price"`
}

type RecipeNameDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
}

func NewRecipeDTO(recipeId int64, recipeName string, recipePrice float64, ingredients []RecipeIngredientDTO) *RecipeDTO {
	return &RecipeDTO{
		Id:          recipeId,
		Name:        recipeName,
		Price:       recipePrice,
		Ingredients: ingredients,
	}
}
