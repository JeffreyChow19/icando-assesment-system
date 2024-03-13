package teacher

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/model/enum"
)

type AuthRoute struct {
	authHandler handler.AuthHandler
}

func (r AuthRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/teacher")

	group.POST(
		"/login", func(c *gin.Context) {
			r.authHandler.Login(c, enum.ROLE_TEACHER)
		},
	)

	group.GET("/profile", r.authHandler.GetTeacherProfile)
}

func NewAuthRoute(authHandler handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		authHandler: authHandler,
	}
}
