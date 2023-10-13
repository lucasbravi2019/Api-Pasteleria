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
	GetPackages(ctx *gin.Context)
	CreatePackage(ctx *gin.Context)
	UpdatePackage(ctx *gin.Context)
	DeletePackage(ctx *gin.Context)
	GetPackageRoutes() []core.Route
}

var PackageHandlerInstance *PackageHandler

func (h *PackageHandler) GetPackages(ctx *gin.Context) {
	statusCode, body := h.Service.GetPackages()
	core.EncodeJsonResponse(ctx.Writer, statusCode, body, nil)
}

func (h *PackageHandler) CreatePackage(ctx *gin.Context) {
	statusCode := h.Service.CreatePackage(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *PackageHandler) UpdatePackage(ctx *gin.Context) {
	statusCode := h.Service.UpdatePackage(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
}

func (h *PackageHandler) DeletePackage(ctx *gin.Context) {
	statusCode := h.Service.DeletePackage(ctx.Request)
	core.EncodeJsonResponse(ctx.Writer, statusCode, nil, nil)
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
