package ingredients

import (
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string              `bson:"name" json:"name,omitempty" validate:"required"`
	Packages []IngredientPackage `bson:"packages" json:"packages,omitempty"`
}

type IngredientPackage struct {
	Package packages.Package `bson:"package" json:"package" validate:"required"`
	Price   float32          `bson:"price" json:"price" validate:"required"`
}

type IngredientPackageDTO struct {
	IngredientOid primitive.ObjectID
	PackageOid    primitive.ObjectID
	Price         float32
}

type IngredientPackagePrice struct {
	Price float32 `json:"price" validate:"required"`
}
