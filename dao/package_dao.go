package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/mapper"
	"github.com/lucasbravi2019/pasteleria/models"
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

<<<<<<< HEAD
func (d *PackageDao) GetPackages() *[]models.Package {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	rows, err := d.DB.Query("")
=======
func (d *PackageDao) GetPackages() (*[]models.Package, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{})

	packages := &[]models.Package{}

	if err != nil {
		log.Println(err.Error())
		return packages, err
	}

	err = cursor.All(ctx, packages)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e

	if err != nil {
		log.Println(err.Error())
		return packages, err
	}

<<<<<<< HEAD
	return mapper.ToPackageList(rows)
}

func (d *PackageDao) CreatePackage(body *models.Package) {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
=======
	return packages, nil
}

func (d *PackageDao) CreatePackage(body *models.Package) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e

	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
<<<<<<< HEAD
	}
=======
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
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

<<<<<<< HEAD
func (d *PackageDao) GetPackageById(oid *primitive.ObjectID) *models.Package {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
=======
func (d *PackageDao) GetPackageById(oid *primitive.ObjectID) (*models.Package, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	defer cancel()

	rows, err := d.DB.Query("")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

<<<<<<< HEAD
	return mapper.ToPackage(rows)
=======
	return envase, nil
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
}
