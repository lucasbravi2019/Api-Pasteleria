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

func (m *PackageMapper) ToPackageList(rows *sql.Rows) (*[]dto.Package, error) {
	packages := util.NewList[dto.Package]()

	for rows.Next() {
		var id int64
		var metric string
		var quantity float64

		err := rows.Scan(&id, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := m.ToPackage(id, metric, quantity)

		if pkg != nil {
			util.Add(&packages, *pkg)
		}
	}

	return &packages, nil
}

func (m *PackageMapper) ToPackage(id int64, metric string, quantity float64) *dto.Package {
	return &dto.Package{
		Id:       &id,
		Metric:   &metric,
		Quantity: &quantity,
	}
}

func (m *PackageMapper) ToPackageNullable(id sql.NullInt64, metric sql.NullString, quantity sql.NullFloat64) *dto.Package {
	if !id.Valid {
		return nil
	}

	return &dto.Package{
		Id:       db.GetLong(id),
		Metric:   db.GetString(metric),
		Quantity: db.GetFloat(quantity),
	}
}
