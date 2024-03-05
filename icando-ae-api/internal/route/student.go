package route

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
)

type StudentRoute struct {
	studentHandler handler.StudentHandler
	//authMiddleware     middleware.AuthMiddleware
}

func (r StudentRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/student")
	group.POST("", r.studentHandler.Post)
	//engine.GET("/protected", r.authMiddleware.Handler(model.ROLE_CUSTOMER), r.healthcheckHandler.StudentProtected)
}

func NewStudentRoute(
	handler handler.StudentHandler,
	// authMiddleware middleware.AuthMiddleware,
) *StudentRoute {
	return &StudentRoute{
		studentHandler: handler,
		//authMiddleware:     authMiddleware,
	}
}
