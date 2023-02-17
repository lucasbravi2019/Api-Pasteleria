package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" validate:"required"`
	Name     string              `bson:"name" validate:"required"`
	Packages []IngredientPackage `bson:"packages" validate:"required"`
}

type IngredientPackage struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Metric   string             `bson:"metric" json:"metric,omitempty"`
	Quantity float64            `bson:"quantity" json:"quantity,omitempty"`
	Price    float64            `bson:"price" json:"price,omitempty"`
}
