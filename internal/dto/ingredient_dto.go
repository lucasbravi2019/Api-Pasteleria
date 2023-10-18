package dto

import (
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type IngredientDTO struct {
	Id       int64        `json:"id,omitempty"`
	Name     string       `json:"name,omitempty" validate:"required"`
	Packages []PackageDTO `json:"packages,omitempty"`
}

type RecipeIngredientDTO struct {
	Id       int64
	Name     string
	Price    float64
	Package  PackageDTO
	Quantity float64
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

func (r *RecipeIngredientDTO) NewRecipeIngredientDTO(id int64, name string, quantity float64, pkg PackageDTO) *RecipeIngredientDTO {
	recipeIngredient := &RecipeIngredientDTO{
		Id:       id,
		Name:     name,
		Quantity: quantity,
		Package:  pkg,
	}

	recipeIngredient.UpdatePrice()
	return recipeIngredient
}

func (r *RecipeIngredientDTO) UpdatePrice() {
	packageQuantity := r.Package.Quantity
	price := r.Package.Price
	recipeQuantity := r.Quantity

	r.Price = recipeQuantity / packageQuantity * price
}
