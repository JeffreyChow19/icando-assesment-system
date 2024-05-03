package teacher

import (
	"icando/internal/handler/teacher"
	"icando/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AnalyticsRoute struct {
	analyticsHandler teacher.AnalyticsHandler
	authMiddleware   middleware.AuthMiddleware
}

func (r AnalyticsRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/analytics")
	group.GET("/overview", r.analyticsHandler.GetDashboardOverview)
	group.GET("/performance", r.analyticsHandler.GetQuizPerformance)
	group.GET("/latest-submissions", r.analyticsHandler.GetLatestSubmissions)
	group.GET("/student/:id", r.analyticsHandler.GetStudentStatistics)
}

func NewAnalyticsRoute(
	handler teacher.AnalyticsHandler,
	authMiddleware *middleware.AuthMiddleware,
) *AnalyticsRoute {
	return &AnalyticsRoute{
		analyticsHandler: handler,
		authMiddleware:   *authMiddleware,
	}
}
