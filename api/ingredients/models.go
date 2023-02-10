package ingredients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       primitive.ObjectID  `bson:"_id" validate:"required"`
	Name     string              `bson:"name" validate:"required"`
	Packages []IngredientPackage `bson:"packages" validate:"required"`
}

type IngredientPackage struct {
	ID    primitive.ObjectID `bson:"_id" validate:"required"`
	Price float64            `bson:"price" validate:"required"`
}
