package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type PackageMapper struct {
}

var PackageMapperInstance *PackageMapper

type PackageMapperInterface interface {
	ToPackageList(rows *sql.Rows) (*[]dto.PackageDTO, error)
	ToIngredientPackage(packageId sql.NullInt64, metric sql.NullString, quantity sql.NullFloat64,
		packagePrice sql.NullFloat64) *dto.IngredientPackageDTO
}

func (m *PackageMapper) ToPackageList(rows *sql.Rows) (*[]dto.PackageDTO, error) {
	packages := util.NewList[dto.PackageDTO]()

	for rows.Next() {
		var id int64
		var metric string
		var quantity float64

		err := rows.Scan(&id, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := m.toPackage(id, metric, quantity)
		util.Add(&packages, pkg)
	}

	return &packages, nil
}

func (m *PackageMapper) ToIngredientPackage(packageId sql.NullInt64, metric sql.NullString, quantity sql.NullFloat64,
	packagePrice sql.NullFloat64) *dto.IngredientPackageDTO {
	if !packageId.Valid {
		return nil
	}

	return &dto.IngredientPackageDTO{
		Id:       db.GetLong(packageId),
		Metric:   db.GetString(metric),
		Quantity: db.GetFloat(quantity),
		Price:    db.GetFloat(packagePrice),
	}
}

func (m *PackageMapper) toPackage(id int64, metric string, quantity float64) dto.PackageDTO {
	return dto.PackageDTO{
		Id:       &id,
		Metric:   &metric,
		Quantity: &quantity,
	}
}
