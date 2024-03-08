package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
	"icando/internal/middleware"
	"icando/internal/model/enum"
)

type ClassRoute struct {
	classHandler   designer.ClassHandler
	authMiddleware middleware.AuthMiddleware
}

func (r ClassRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/class")
	group.GET("/", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.classHandler.GetAll)
	group.POST("", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.classHandler.Create)
	group.GET("/:id", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.classHandler.Get)
	group.PATCH("/:id", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.classHandler.Update)
	group.DELETE("/:id", r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER), r.classHandler.Delete)
}

func NewClassRoute(
	classHandler designer.ClassHandler,
	authMiddleware *middleware.AuthMiddleware,
) *ClassRoute {
	return &ClassRoute{
		classHandler:   classHandler,
		authMiddleware: *authMiddleware,
	}
}
