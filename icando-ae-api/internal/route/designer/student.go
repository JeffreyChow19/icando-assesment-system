package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
	"icando/internal/middleware"
)

type StudentRoute struct {
	studentHandler designer.StudentHandler
	authMiddleware middleware.AuthMiddleware
}

func (r StudentRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/student")
	group.POST("", r.studentHandler.Post)
	group.GET("/:id", r.studentHandler.Get)
	group.PATCH("/:id", r.studentHandler.Patch)
	group.DELETE("/:id", r.studentHandler.Delete)
	group.GET("", r.studentHandler.GetAll)
}

func NewStudentRoute(
	handler designer.StudentHandler,
	authMiddleware *middleware.AuthMiddleware,
) *StudentRoute {
	return &StudentRoute{
		studentHandler: handler,
		authMiddleware: *authMiddleware,
	}
}
