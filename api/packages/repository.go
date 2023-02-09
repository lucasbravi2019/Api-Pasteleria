package packages

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Collection
}

type PackageRepository interface {
	GetPackages() (int, []Package)
	CreatePackage(body *Package) (int, *Package)
	UpdatePackage(oid *primitive.ObjectID, body *Package) (int, *Package)
	DeletePackage(oid *primitive.ObjectID) (int, *Package)
}

var packageRepositoryInstance *repository

func (r *repository) GetPackages() (int, []Package) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	cursor, err := r.db.Find(ctx, bson.M{})

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var packages []Package

	err = cursor.All(ctx, &packages)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	if packages == nil {
		return http.StatusOK, []Package{}
	}

	return http.StatusOK, packages
}

func (r *repository) CreatePackage(body *Package) (int, *Package) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	result, err := r.db.InsertOne(ctx, body)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	if result.InsertedID == nil {
		return http.StatusInternalServerError, nil
	}

	body.ID = result.InsertedID.(primitive.ObjectID)

	return http.StatusCreated, body
}

func (r *repository) UpdatePackage(oid *primitive.ObjectID, body *Package) (int, *Package) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	err := r.db.FindOneAndUpdate(ctx, GetPackageById(*oid), UpdatePackageById(*body)).Decode(body)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, body
}

func (r *repository) DeletePackage(oid *primitive.ObjectID) (int, *Package) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)

	defer cancel()

	var packageDeleted *Package = &Package{}

	err := r.db.FindOneAndDelete(ctx, GetPackageById(*oid)).Decode(packageDeleted)

	if err != nil {
		log.Println("No pudo borrarse el paquete")
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packageDeleted
}
