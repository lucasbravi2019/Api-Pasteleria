package packages

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Collection
}

type PackageRepository interface {
	GetPackages() *[]Package
	GetPackageById(oid *primitive.ObjectID) *Package
	CreatePackage(body *Package) *primitive.ObjectID
	UpdatePackage(oid *primitive.ObjectID, body *Package) error
	DeletePackage(oid *primitive.ObjectID) error
}

var packageRepositoryInstance *repository

func (r *repository) GetPackages() *[]Package {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	cursor, err := r.db.Find(ctx, bson.M{})

	var packages *[]Package = &[]Package{}

	if err != nil {
		log.Println(err.Error())
		return packages
	}

	err = cursor.All(ctx, packages)

	if err != nil {
		log.Println(err.Error())
	}

	return packages
}

func (r *repository) CreatePackage(body *Package) *primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	result, err := r.db.InsertOne(ctx, body)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if result.InsertedID == nil {
		return nil
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id
}

func (r *repository) UpdatePackage(oid *primitive.ObjectID, body *Package) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetPackageById(*oid), UpdatePackageById(*body))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) DeletePackage(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	_, err := r.db.DeleteOne(ctx, GetPackageById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) GetPackageById(oid *primitive.ObjectID) *Package {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var envase *Package = &Package{}

	err := r.db.FindOne(ctx, GetPackageById(*oid)).Decode(envase)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return envase
}
