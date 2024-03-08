package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/model"
)

type StudentRoute struct {
	studentHandler handler.StudentHandler
	authMiddleware middleware.AuthMiddleware
}

func (r StudentRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/student")
	group.POST("", r.authMiddleware.Handler(model.ROLE_LEARNING_DESIGNER), r.studentHandler.Post)
	group.GET("/:id", r.authMiddleware.Handler(model.ROLE_LEARNING_DESIGNER), r.studentHandler.Get)
	group.PATCH("/:id", r.authMiddleware.Handler(model.ROLE_LEARNING_DESIGNER), r.studentHandler.Patch)
	group.DELETE("/:id", r.authMiddleware.Handler(model.ROLE_LEARNING_DESIGNER), r.studentHandler.Delete)
}

func NewStudentRoute(
	handler handler.StudentHandler,
	authMiddleware middleware.AuthMiddleware,
) *StudentRoute {
	return &StudentRoute{
		studentHandler: handler,
		authMiddleware: authMiddleware,
	}
}
