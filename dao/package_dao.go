package dao

import (
	"context"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/models"
	"github.com/lucasbravi2019/pasteleria/queries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PackageDao struct {
	DB *mongo.Collection
}

type PackageDaoInterface interface {
	GetPackages() (*[]models.Package, error)
	GetPackageById(oid *primitive.ObjectID) (*models.Package, error)
	CreatePackage(body *models.Package) (*primitive.ObjectID, error)
	UpdatePackage(oid *primitive.ObjectID, body *models.Package) error
	DeletePackage(oid *primitive.ObjectID) error
}

var PackageDaoInstance *PackageDao

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

	if err != nil {
		log.Println(err.Error())
		return packages, err
	}

	return packages, nil
}

func (d *PackageDao) CreatePackage(body *models.Package) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	result, err := d.DB.InsertOne(ctx, body)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
}

func (d *PackageDao) UpdatePackage(oid *primitive.ObjectID, body *models.Package) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetPackageById(*oid), queries.UpdatePackageById(*body))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *PackageDao) DeletePackage(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := d.DB.DeleteOne(ctx, queries.GetPackageById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *PackageDao) GetPackageById(oid *primitive.ObjectID) (*models.Package, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	envase := &models.Package{}

	err := d.DB.FindOne(ctx, queries.GetPackageById(*oid)).Decode(envase)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return envase, nil
}
