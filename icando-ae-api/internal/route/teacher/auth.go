package teacher

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/model"
)

type AuthRoute struct {
	authHandler    handler.AuthHandler
	authMiddleware middleware.AuthMiddleware
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/auth")

	group.POST("/login", func(c *gin.Context) {
		r.authHandler.Login(c, model.ROLE_TEACHER)
	})
}

func NewAuthRoute(authHandler handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		authHandler: authHandler,
	}
}
