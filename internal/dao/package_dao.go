package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PackageDao struct {
	DB *sql.DB
}

type PackageDaoInterface interface {
	GetPackages() (*[]models.Package, error)
	GetPackageById(oid *primitive.ObjectID) (*models.Package, error)
	CreatePackage(body *models.Package) (*primitive.ObjectID, error)
	UpdatePackage(oid *primitive.ObjectID, body *models.Package) error
	DeletePackage(oid *primitive.ObjectID) error
}

var PackageDaoInstance *PackageDao

func (d *PackageDao) GetPackages() *[]models.Package {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return mapper.ToPackageList(rows)
}

func (d *PackageDao) CreatePackage(body *models.Package) {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}
}

func (d *PackageDao) UpdatePackage(oid *primitive.ObjectID, body *models.Package) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *PackageDao) DeletePackage(oid *primitive.ObjectID) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *PackageDao) GetPackageById(oid *primitive.ObjectID) *models.Package {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return mapper.ToPackage(rows)
}
