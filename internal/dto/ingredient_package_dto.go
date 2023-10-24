package dto

import "github.com/lucasbravi2019/pasteleria/pkg/util"

type IngredientPackageId struct {
	IngredientId int64 `json:"ingredientId" validate:"required"`
}

type IngredientPackagePrices struct {
	IngredientId int64          `json:"ingredientId" validate:"required"`
	Packages     []PackagePrice `json:"packages" validate:"required"`
}

type PackagePrice struct {
	PackageId int64   `json:"packageId" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}

func (i *IngredientPackagePrices) AddPackage(packageId int64, price float64) {
	pkg := PackagePrice{
		PackageId: packageId,
		Price:     price,
	}

	util.Add(&i.Packages, pkg)
}

func (i *IngredientPackagePrices) GetPackagesQuantity() int {
	if i.Packages == nil {
		return 0
	}

	return len(i.Packages)
}
