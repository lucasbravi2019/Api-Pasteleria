package models

import "github.com/lucasbravi2019/pasteleria/pkg/util"

type Recipe struct {
	Id          int64
	Name        string
	Ingredients []RecipeIngredient
	Price       float64
}

type RecipeIngredient struct {
	IngredientId *int64
	Name         *string
	RecipeId     *int64
	Package      *RecipeIngredientPackage
	Quantity     *float64
	Price        *float64
}

type RecipeIngredientPackage struct {
	PackageId int64
	Metric    string
	Quantity  float64
	Price     float64
}

func NewRecipe(recipeId int64, recipeName string, recipePrice float64) *Recipe {
	return &Recipe{
		Id:          recipeId,
		Name:        recipeName,
		Price:       recipePrice,
		Ingredients: util.NewList[RecipeIngredient](),
	}
}

func NewRecipeIngredientPackage(packageId *int64, metric *string, quantity *float64, price *float64) *RecipeIngredientPackage {
	if packageId == nil {
		return nil
	}
	return &RecipeIngredientPackage{
		PackageId: *packageId,
		Metric:    *metric,
		Quantity:  *quantity,
		Price:     *price,
	}
}

func NewRecipeIngredient(ingredientId *int64, ingredientName *string, quantity *float64, price *float64, recipeId *int64) *RecipeIngredient {
	return &RecipeIngredient{
		IngredientId: ingredientId,
		Name:         ingredientName,
		Quantity:     quantity,
		Price:        price,
		RecipeId:     recipeId,
	}
}
