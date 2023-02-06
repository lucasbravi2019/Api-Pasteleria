package recipes

import (
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty" validate:"required"`
	Ingredients []RecipeIngredient `bson:"ingredients" json:"ingredients,omitempty" validate:"required"`
	Price       float32            `bson:"price" json:"price,omitempty" validate:"required"`
}

type RecipeIngredient struct {
	ID       primitive.ObjectID            `bson:"_id" json:"id,omitempty"`
	Name     string                        `bson:"name" json:"name" validate:"required"`
	Package  ingredients.IngredientPackage `bson:"package" json:"package" validate:"required"`
	Price    float32                       `bson:"price" json:"price,omitempty"`
	Quantity int                           `bson:"quantity" json:"quantity,omitempty"`
}

type IngredientDetails struct {
	Metric   string `json:"metric,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type RecipeName struct {
	Name string `json:"name" validate:"required"`
}
