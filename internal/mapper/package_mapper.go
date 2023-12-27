package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type PackageMapper struct {
}

var PackageMapperInstance *PackageMapper

type PackageMapperInterface interface {
	ToPackageList(rows *sql.Rows) (*[]models.Package, error)

	ToPackageDTO(pkg *models.Package) *dto.PackageDTO

	ToPackageDTOList(pkgs *[]models.Package) *[]dto.PackageDTO

	ToRecipeIngredientPackage(packageId sql.NullInt64, metric sql.NullString, quantity sql.NullFloat64,
		packagePrice sql.NullFloat64) *models.RecipeIngredientPackage
}

func (m *PackageMapper) ToPackageList(rows *sql.Rows) (*[]models.Package, error) {
	packages := util.NewList[models.Package]()

	for rows.Next() {
		var id int64
		var metric string
		var quantity float64

		err := rows.Scan(&id, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := models.NewPackage(id, metric, quantity)

		util.Add(&packages, *pkg)
	}

	return &packages, nil
}

func (m *PackageMapper) ToPackageDTO(pkg *models.Package) *dto.PackageDTO {
	return dto.NewPackageDTO(pkg.Id, pkg.Metric, pkg.Quantity, 0)
}

func (m *PackageMapper) ToPackageDTOList(pkgs *[]models.Package) *[]dto.PackageDTO {
	dtos := util.NewList[dto.PackageDTO]()

	for _, pkg := range *pkgs {
		dto := dto.NewPackageDTO(pkg.Id, pkg.Metric, pkg.Quantity, 0)

		util.Add(&dtos, *dto)
	}

	return &dtos
}

func (m *PackageMapper) ToRecipeIngredientPackage(packageId sql.NullInt64, metric sql.NullString,
	quantity sql.NullFloat64, packagePrice sql.NullFloat64) *models.RecipeIngredientPackage {
	return models.NewRecipeIngredientPackage(db.GetLong(packageId), db.GetString(metric), db.GetFloat(quantity), db.GetFloat(packagePrice))
}
