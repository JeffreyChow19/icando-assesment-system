package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
)

type ClassRoute struct {
	classHandler designer.ClassHandler
}

func (r ClassRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/class")
	group.GET("/", r.classHandler.GetAll)
	group.POST("", r.classHandler.Create)
	// todo: remove redundant route /:id/students
	group.GET("/:id/students", r.classHandler.GetWithStudents)
	group.POST("/:id/students", r.classHandler.AssignStudents)
	group.PATCH("/:id/students", r.classHandler.UnassignStudents)
	group.GET("/:id", r.classHandler.Get)
	group.PATCH("/:id", r.classHandler.Update)
	group.DELETE("/:id", r.classHandler.Delete)
}

func NewClassRoute(
	classHandler designer.ClassHandler,
) *ClassRoute {
	return &ClassRoute{
		classHandler: classHandler,
	}
}
