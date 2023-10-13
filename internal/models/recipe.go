package models

type Recipe struct {
	ID          int
	Name        string
	Ingredients []RecipeIngredient
	Price       float64
}

type RecipeIngredient struct {
	ID       int
	Name     string
	Package  RecipeIngredientPackage
	Quantity float64
	Price    float64
}

type RecipeIngredientPackage struct {
	ID       int
	Metric   string
	Quantity float64
	Price    float64
}
