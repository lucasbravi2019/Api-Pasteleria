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

type PackageDTO struct {
	ID       int     `json:"id,omitempty"`
	Metric   string  `json:"metric,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

type RecipeNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type RecipeIdDTO struct {
	ID string `json:"id" validate:"required"`
}
