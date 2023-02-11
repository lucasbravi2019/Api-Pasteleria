package recipes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty" validate:"required"`
	Ingredients []RecipeIngredient `bson:"ingredients" json:"ingredients,omitempty" validate:"required"`
	Price       float64            `bson:"price" json:"price,omitempty" validate:"required"`
}

type RecipeIngredient struct {
	ID       primitive.ObjectID      `bson:"_id" json:"id,omitempty"`
	Name     string                  `bson:"name" json:"name,omitempty"`
	Package  RecipeIngredientPackage `bson:"package" json:"package" validate:"required"`
	Quantity float32                 `bson:"quantity" json:"quantity" validate:"required"`
	Price    float64                 `bson:"price" json:"price" validate:"required"`
}

type RecipeIngredientPackage struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" validate:"required"`
	Metric   string             `bson:"metric" json:"metric"`
	Quantity float64            `bson:"quantity" json:"quantity"`
	Price    float64            `bson:"price" json:"price"`
}
