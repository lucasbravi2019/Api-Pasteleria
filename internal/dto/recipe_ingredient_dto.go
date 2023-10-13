package dto

type RecipeIngredientIdDTO struct {
	RecipeId     string `bson:"_id" json:"recipeId" validate:"required"`
	IngredientId string `bson:"_id" json:"ingredientId" validate:"required"`
}
