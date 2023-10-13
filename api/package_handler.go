package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/services"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type PackageHandler struct {
	Service services.PackageServiceInterface
}

type PackageHandlerInterface interface {
	GetPackages(ctx *gin.Context)
	CreatePackage(ctx *gin.Context)
	UpdatePackage(ctx *gin.Context)
	DeletePackage(ctx *gin.Context)
	GetPackageRoutes() []pkg.Route
}

var PackageHandlerInstance *PackageHandler

func (h *PackageHandler) GetPackages(ctx *gin.Context) {
	statusCode, body, err := h.Service.GetPackages()
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *PackageHandler) CreatePackage(ctx *gin.Context) {
	statusCode, body, err := h.Service.CreatePackage(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *PackageHandler) UpdatePackage(ctx *gin.Context) {
	statusCode, body, err := h.Service.UpdatePackage(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *PackageHandler) DeletePackage(ctx *gin.Context) {
	statusCode, body, err := h.Service.DeletePackage(ctx)
	pkg.EncodeJsonResponse(ctx, statusCode, body, err)
}

func (h *PackageHandler) GetPackageRoutes() pkg.Routes {
	return pkg.Routes{
		pkg.Route{
			Path:        "/packages",
			HandlerFunc: h.GetPackages,
			Method:      "GET",
		},
		pkg.Route{
			Path:        "/packages",
			HandlerFunc: h.CreatePackage,
			Method:      "POST",
		},
		pkg.Route{
			Path:        "/packages/:packageId",
			HandlerFunc: h.UpdatePackage,
			Method:      "PUT",
		},
		pkg.Route{
			Path:        "/packages/:packageId",
			HandlerFunc: h.DeletePackage,
			Method:      "DELETE",
		},
	}
}
