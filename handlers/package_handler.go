package handlers

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type PackageHandler struct {
	Service services.PackageServiceInterface
}

type PackageHandlerInterface interface {
	GetPackages(w http.ResponseWriter, r *http.Request)
	CreatePackage(w http.ResponseWriter, r *http.Request)
	UpdatePackage(w http.ResponseWriter, r *http.Request)
	DeletePackage(w http.ResponseWriter, r *http.Request)
	GetPackageRoutes() []core.Route
}

var PackageHandlerInstance *PackageHandler

func (h *PackageHandler) GetPackages(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.GetPackages()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *PackageHandler) CreatePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.CreatePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *PackageHandler) UpdatePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.UpdatePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *PackageHandler) DeletePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.Service.DeletePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *PackageHandler) GetPackageRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/packages",
			HandlerFunc: h.GetPackages,
			Method:      "GET",
		},
		core.Route{
			Path:        "/packages",
			HandlerFunc: h.CreatePackage,
			Method:      "POST",
		},
		core.Route{
			Path:        "/packages/{id}",
			HandlerFunc: h.UpdatePackage,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/packages/{id}",
			HandlerFunc: h.DeletePackage,
			Method:      "DELETE",
		},
	}
}
