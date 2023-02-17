package dto

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
	Package  PackageDTO         `json:"package"`
	Quantity float64            `json:"quantity"`
}

type PackageDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Metric   string             `bson:"metric,omitempty" json:"metric,omitempty"`
	Quantity float64            `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Price    float64            `bson:"price,omitempty" json:"price,omitempty"`
}

type RecipeNameDTO struct {
	Name string `json:"name" validate:"required"`
}
