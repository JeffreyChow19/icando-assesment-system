package route

import (
	"icando/internal/handler"
	"icando/internal/middleware"
	// "icando/internal/model"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authHandler    handler.AuthHandler
	authMiddleware middleware.AuthMiddleware
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/auth")
	group.POST("/login", r.authMiddleware.Handler("Learning Designer"), r.authHandler.Login)
}

func NewAuthRoute(authHandler handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		authHandler: authHandler,
	}
}
