package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/model/enum"
)

type AuthRoute struct {
	authHandler    handler.AuthHandler
	authMiddleware middleware.AuthMiddleware
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/auth")
	group.POST("/login", func(c *gin.Context) {
		r.authHandler.Login(c, enum.ROLE_LEARNING_DESIGNER)
	})
	group.GET("/profile", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.authHandler.GetLearningDesignerProfile)
}

func NewAuthRoute(authHandler handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) *AuthRoute {
	return &AuthRoute{
		authHandler:    authHandler,
		authMiddleware: *authMiddleware,
	}
}
