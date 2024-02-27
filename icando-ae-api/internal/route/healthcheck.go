package route

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
)

type HealthcheckRoute struct {
	healthcheckHandler handler.HealthcheckHandler
	//authMiddleware     middleware.AuthMiddleware
}

func (r HealthcheckRoute) Setup(engine *gin.Engine) {
	engine.GET("/", r.healthcheckHandler.Healthcheck)
	//engine.GET("/protected", r.authMiddleware.Handler(model.ROLE_CUSTOMER), r.healthcheckHandler.HealthcheckProtected)
}

func NewHealthcheckRoute(
	handler handler.HealthcheckHandler,
	// authMiddleware middleware.AuthMiddleware,
) *HealthcheckRoute {
	return &HealthcheckRoute{
		healthcheckHandler: handler,
		//authMiddleware:     authMiddleware,
	}
}
