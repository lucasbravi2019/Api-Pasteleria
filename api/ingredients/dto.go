package ingredients

import (
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientDTO struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name" json:"name,omitempty" validate:"required"`
	Package []packages.Package `bson:"package" json:"package,omitempty"`
}

type IngredientPackageDTO struct {
	IngredientOid primitive.ObjectID
	PackageOid    primitive.ObjectID
	Price         float64
}

type IngredientPackagePriceDTO struct {
	Price float64 `json:"price" validate:"required"`
}
