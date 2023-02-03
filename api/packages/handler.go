package packages

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/core"
)

type handler struct {
	service PackageService
}

type PackageHandler interface {
	GetPackages(w http.ResponseWriter, r *http.Request)
	CreatePackage(w http.ResponseWriter, r *http.Request)
	UpdatePackage(w http.ResponseWriter, r *http.Request)
	DeletePackage(w http.ResponseWriter, r *http.Request)
	GetPackageRoutes() []core.Route
}

var packageHandlerInstance *handler

func (h *handler) GetPackages(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.GetPackages()
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) CreatePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.CreatePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) UpdatePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.UpdatePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) DeletePackage(w http.ResponseWriter, r *http.Request) {
	statusCode, body := h.service.DeletePackage(r)
	core.EncodeJsonResponse(w, statusCode, body)
}

func (h *handler) GetPackageRoutes() core.Routes {
	return core.Routes{
		core.Route{
			Path:        "/package",
			HandlerFunc: h.GetPackages,
			Method:      "GET",
		},
		core.Route{
			Path:        "/package",
			HandlerFunc: h.CreatePackage,
			Method:      "POST",
		},
		core.Route{
			Path:        "/package/{id}",
			HandlerFunc: h.UpdatePackage,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/package/{id}",
			HandlerFunc: h.DeletePackage,
			Method:      "DELETE",
		},
	}
}
