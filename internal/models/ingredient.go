package models

import "github.com/lucasbravi2019/pasteleria/pkg/util"

type Ingredient struct {
	Id       int64
	Name     string
	Packages []IngredientPackage
}

type IngredientPackage struct {
	Id       int64
	Metric   string
	Quantity float64
	Price    float64
}

func NewIngredient(id int64, name string) *Ingredient {
	return &Ingredient{
		Id:       id,
		Name:     name,
		Packages: util.NewList[IngredientPackage](),
	}
}

func (i *Ingredient) AddPackage(p *IngredientPackage) {
	if p != nil {
		util.Add(&i.Packages, *p)
	}
}
