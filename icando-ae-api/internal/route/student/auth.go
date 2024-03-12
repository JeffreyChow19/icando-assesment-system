package student

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
)

type AuthRoute struct {
	authHandler handler.AuthHandler
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/student")

	group.GET("/profile", r.authHandler.GetStudentProfile)
}

func NewAuthRoute(authHandler handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		authHandler: authHandler,
	}
}
