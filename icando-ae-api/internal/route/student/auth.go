package student

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/model/enum"
)

type AuthRoute struct {
	authHandler    handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/student")

	group.GET("/profile", r.authMiddleware.Handler(enum.ROLE_STUDENT), r.authHandler.GetStudentProfile)
}

func NewAuthRoute(authHandler handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) *AuthRoute {
	return &AuthRoute{
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
	}
}
