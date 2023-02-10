package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeDTO struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name"`
	Ingredients []IngredientsDTO   `json:"ingredients"`
	Price       float64            `json:"price"`
}

type IngredientsDTO struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	Price    float64            `json:"price"`
	Packages PackageDTO         `json:"packages"`
	Quantity float64            `json:"quantity"`
}

type PackageDTO struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Metric   string             `json:"metric"`
	Quantity float64            `json:"quantity"`
}

type IngredientDetailsDTO struct {
	Metric   string  `json:"metric,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}

type RecipeNameDTO struct {
	Name string `json:"name" validate:"required"`
}
