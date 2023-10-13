package models

type Ingredient struct {
	ID       int
	Name     string
	Packages []IngredientPackage
}

type IngredientPackage struct {
	ID       int
	Metric   string
	Quantity float64
	Price    float64
}
