package dto

import (
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientCreationDTO struct {
	Name     string                      `json:"name" validate:"required"`
	Packages []IngredientPackagePriceDTO `json:"packages"`
}

type IngredientUpdateDTO struct {
	Id       int64                       `json:"id"`
	Name     string                      `json:"name" validate:"required"`
	Packages []IngredientPackagePriceDTO `json:"packages"`
}

type IngredientPackagePriceDTO struct {
	Id    int64   `json:"id"`
	Price float64 `json:"price"`
}

type IngredientPackageDTO struct {
	Id       *int64   `json:"id"`
	Metric   *string  `json:"metric"`
	Quantity *float64 `json:"quantity"`
	Price    *float64 `json:"price"`
}

type IngredientDTO struct {
	Id       int64        `json:"id,omitempty"`
	Name     string       `json:"name,omitempty" validate:"required"`
	Packages []PackageDTO `json:"packages,omitempty"`
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
