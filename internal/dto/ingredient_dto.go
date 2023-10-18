package dto

import (
	"github.com/lucasbravi2019/pasteleria/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type IngredientDTO struct {
	Id       int64        `json:"id,omitempty"`
	Name     string       `json:"name,omitempty" validate:"required"`
	Packages []PackageDTO `json:"packages,omitempty"`
}

type IngredientPackagePriceDTO struct {
	Price float64 `json:"price" validate:"required"`
}

type RecipeIngredientDTO struct {
	ID       primitive.ObjectID
	Name     string
	Price    float64
	Package  PackageDTO
	Quantity float64
}

type IngredientDetailsDTO struct {
	Metric   string  `json:"metric,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}

func NewIngredientDTO(id int64, name string) *IngredientDTO {
	return &IngredientDTO{
		Id:       id,
		Name:     name,
		Packages: util.NewList[PackageDTO](),
	}
}

func (i *IngredientDTO) AddPackage(pkg *PackageDTO) {
	if pkg != nil {
		util.Add(&i.Packages, *pkg)
	}
}
