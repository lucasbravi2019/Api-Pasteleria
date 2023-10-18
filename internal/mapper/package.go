package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

func ToPackageList(rows *sql.Rows) (*[]models.Package, error) {
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

func ToPackageDTO(pkg *models.Package) *dto.PackageDTO {
	return dto.NewPackageDTO(pkg.Id, pkg.Metric, pkg.Quantity)
}

func ToPackageDTOList(pkgs *[]models.Package) *[]dto.PackageDTO {
	dtos := util.NewList[dto.PackageDTO]()

	for _, pkg := range *pkgs {
		dto := dto.NewPackageDTO(pkg.Id, pkg.Metric, pkg.Quantity)

		util.Add(&dtos, *dto)
	}

	return &dtos
}
