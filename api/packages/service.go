package packages

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
)

type service struct {
	repository PackageRepository
}

type PackageService interface {
	GetPackages() (int, []Package)
	CreatePackage(r *http.Request) (int, *Package)
	UpdatePackage(r *http.Request) (int, *Package)
	DeletePackage(r *http.Request) (int, *Package)
}

var packageServiceInstance *service

func (s *service) GetPackages() (int, []Package) {
	return s.repository.GetPackages()
}

func (s *service) CreatePackage(r *http.Request) (int, *Package) {
	var packageRequest *Package = &Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}
	return s.repository.CreatePackage(packageRequest)
}

func (s *service) UpdatePackage(r *http.Request) (int, *Package) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var packageRequest *Package = &Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.UpdatePackage(oid, packageRequest)
}

func (s *service) DeletePackage(r *http.Request) (int, *Package) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.repository.DeletePackage(oid)
}
