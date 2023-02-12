package ingredients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type IngredientDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Packages []PackageDTO       `bson:"packages,omitempty" json:"packages,omitempty"`
}

type PackageDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Metric   string             `bson:"metric,omitempty" json:"metric,omitempty"`
	Quantity float64            `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Price    float64            `bson:"price,omitempty" json:"price,omitempty"`
}

type IngredientPackageDTO struct {
	IngredientOid primitive.ObjectID
	PackageOid    primitive.ObjectID
	Price         float64
}

type IngredientPackagePriceDTO struct {
	Price float64 `json:"price" validate:"required"`
}

type RecipeIngredientDTO struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	Price    float64            `json:"price"`
	Package  PackageDTO         `json:"package"`
	Quantity float64            `json:"quantity"`
}

type IngredientDetailsDTO struct {
	Metric   string  `json:"metric,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}
