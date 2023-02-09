package ingredients

import (
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty" validate:"required"`
	Packages []packages.Package `bson:"packages" json:"packages,omitempty"`
}

type IngredientPackage struct {
	Package  packages.Package `bson:"package" json:"package" validate:"required"`
	Price    float64          `bson:"price" json:"price" validate:"required"`
	Quantity float32          `bson:"quantity" json:"quantity"`
}

type IngredientPackageDTO struct {
	IngredientOid primitive.ObjectID
	PackageOid    primitive.ObjectID
	Price         float64
}

type IngredientPackagePrice struct {
	Price float64 `json:"price" validate:"required"`
}
