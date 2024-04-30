package teacher

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
	"icando/internal/middleware"
)

type AnalyticsRoute struct {
	teacherHandler designer.TeacherHandler
	authMiddleware middleware.AuthMiddleware
}

func (r AnalyticsRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/analytics")
	group.GET("/dashboard/overview", r.teacherHandler.GetDashboardOverview)
}

func NewAnalyticsRoute(
	teacherHandler designer.TeacherHandler,
	authMiddleware *middleware.AuthMiddleware,
) *AnalyticsRoute {
	return &AnalyticsRoute{
		teacherHandler: teacherHandler,
		authMiddleware: *authMiddleware,
	}
}
