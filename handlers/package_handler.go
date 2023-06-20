package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/services"
)

type PackageHandler struct {
	Service services.PackageServiceInterface
}

type PackageHandlerInterface interface {
	GetPackages(c *gin.Context)
	CreatePackage(c *gin.Context)
	UpdatePackage(c *gin.Context)
	DeletePackage(c *gin.Context)
	GetPackageRoutes() []core.Route
}

var PackageHandlerInstance *PackageHandler

func (h *PackageHandler) GetPackages(c *gin.Context) {
	statusCode, body, err := h.Service.GetPackages()
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *PackageHandler) CreatePackage(c *gin.Context) {
	statusCode, body, err := h.Service.CreatePackage(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	statusCode, body, err := h.Service.UpdatePackage(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	statusCode, body, err := h.Service.DeletePackage(c)
	core.EncodeJsonResponse(c, statusCode, body, err)
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
			Path:        "/packages/:packageId",
			HandlerFunc: h.UpdatePackage,
			Method:      "PUT",
		},
		core.Route{
			Path:        "/packages/:packageId",
			HandlerFunc: h.DeletePackage,
			Method:      "DELETE",
		},
	}
}
