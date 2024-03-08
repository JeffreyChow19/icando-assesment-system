package route

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/middleware"
)

type HealthcheckRoute struct {
	healthcheckHandler handler.HealthcheckHandler
	authMiddleware     middleware.AuthMiddleware
}

func (r HealthcheckRoute) Setup(group *gin.RouterGroup) {
	group.GET("/", r.healthcheckHandler.Healthcheck)
	group.GET("/protected", r.authMiddleware.Handler("Learning Designer"), r.healthcheckHandler.HealthcheckProtected)
}

func NewHealthcheckRoute(
	handler handler.HealthcheckHandler,
	authMiddleware *middleware.AuthMiddleware,
) *HealthcheckRoute {
	return &HealthcheckRoute{
		healthcheckHandler: handler,
		authMiddleware:     *authMiddleware,
	}
}
