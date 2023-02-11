package ingredients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type IngredientDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty" validate:"required"`
	Packages []PackageDTO       `bson:"packages" json:"packages,omitempty"`
}

type PackageDTO struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Metric   string             `bson:"metric" json:"metric,omitempty"`
	Quantity float64            `bson:"quantity" json:"quantity,omitempty"`
	Price    float64            `bson:"price" json:"price,omitempty"`
}

type IngredientPackageDTO struct {
	IngredientOid primitive.ObjectID
	PackageOid    primitive.ObjectID
	Price         float64
}

type IngredientPackagePriceDTO struct {
	Price float64 `json:"price" validate:"required"`
}
