package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
)

type TeacherRoute struct {
	teacherHandler designer.TeacherHandler
}

func (r TeacherRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/teacher")
	group.GET("", r.teacherHandler.GetAll)
}

func NewTeacherRoute(
	handler designer.TeacherHandler,
) *TeacherRoute {
	return &TeacherRoute{
		teacherHandler: handler,
	}
}
