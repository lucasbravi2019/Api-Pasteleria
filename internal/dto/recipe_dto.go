package dto

type RecipeDTO struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Ingredients []IngredientsDTO `json:"ingredients"`
	Price       float64          `json:"price"`
}

type IngredientsDTO struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	Package  PackageDTO `json:"package"`
	Quantity float64    `json:"quantity"`
}

type RecipeNameDTO struct {
	Name string `json:"name" validate:"required"`
}
