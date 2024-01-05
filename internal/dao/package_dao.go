package dao

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type PackageDao struct {
	DB            *sql.DB
	PackageMapper *mapper.PackageMapper
}

type PackageDaoInterface interface {
	GetPackages() (*[]dto.PackageDTO, error)
	CreatePackage(body *dto.PackageDTO) (*int64, error)
	UpdatePackage(body *dto.PackageDTO) error
	DeletePackage(id *int64) error
}

var PackageDaoInstance *PackageDao

func (d *PackageDao) GetPackages() (*[]dto.PackageDTO, error) {
	query, err := db.GetQueryByName(db.Package_FindAll)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query)

	if pkg.HasError(err) {
		return nil, err
	}

	return d.PackageMapper.ToPackageList(rows)
}

func (d *PackageDao) CreatePackage(body *dto.PackageDTO) error {
	query, err := db.GetQueryByName(db.Package_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, body.Metric, body.Quantity)

	return err
}

func (d *PackageDao) UpdatePackage(body *dto.PackageDTO) error {
	query, err := db.GetQueryByName(db.Package_UpdateById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, body.Metric, body.Quantity, body.Id)

	return err
}

func (d *PackageDao) DeletePackage(id *int64) error {
	query, err := db.GetQueryByName(db.Package_DeleteById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, id)

	return err
}
