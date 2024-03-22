package designer

import (
	"icando/internal/handler/designer"
	"icando/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ClassRoute struct {
	classHandler designer.ClassHandler
	authMiddleware middleware.AuthMiddleware
}

func (r ClassRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/class")
	group.GET("", r.classHandler.GetAll)
	group.POST("", r.classHandler.Create)
	group.POST("/:id/students", r.classHandler.AssignStudents)
	group.PATCH("/:id/students", r.classHandler.UnassignStudents)
	group.GET("/:id", r.classHandler.Get)
	group.PATCH("/:id", r.classHandler.Update)
	group.DELETE("/:id", r.classHandler.Delete)
}

func NewClassRoute(
	classHandler designer.ClassHandler,
	authMiddleware *middleware.AuthMiddleware,
) *ClassRoute {
	return &ClassRoute{
		classHandler: classHandler,
		authMiddleware: *authMiddleware,
	}
}
