package dto

import (
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientNameDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
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
