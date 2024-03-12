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
