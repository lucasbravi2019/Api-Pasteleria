package recipes

import (
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
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
	Package  RecipeIngredientPackage `bson:"package" json:"package" validate:"required"`
	Quantity float32                 `bson:"quantity" json:"quantity" validate:"required"`
}

type RecipeIngredientPackage struct {
	ID      primitive.ObjectID            `bson:"_id,omitempty" validate:"required"`
	Name    string                        `bson:"name" validate:"required"`
	Package ingredients.IngredientPackage `bson:"package" validate:"required"`
}
