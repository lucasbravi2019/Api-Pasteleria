package models

type Recipe struct {
	Id          int64
	Name        string
	Ingredients []RecipeIngredient
	Price       *float64
}

type RecipeIngredient struct {
	IngredientId int64
	Name         string
	Package      *RecipeIngredientPackage
	Quantity     float64
	Price        *float64
}

type RecipeIngredientPackage struct {
	PackageId int64
	Metric    string
	Quantity  float64
	Price     float64
}

func NewRecipe(recipeId int64, recipeName string, recipeIngredients []RecipeIngredient, recipePrice *float64) *Recipe {
	return &Recipe{
		Id:          recipeId,
		Name:        recipeName,
		Ingredients: recipeIngredients,
		Price:       recipePrice,
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

func NewRecipeIngredient(ingredientId int64, ingredientName string, pkg *RecipeIngredientPackage, quantity float64, price *float64) *RecipeIngredient {
	return &RecipeIngredient{
		IngredientId: ingredientId,
		Name:         ingredientName,
		Package:      pkg,
		Quantity:     quantity,
		Price:        price,
	}
}
