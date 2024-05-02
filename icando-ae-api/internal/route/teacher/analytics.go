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
	group.GET("/performance", r.analyticsHandler.GetQuizPerformance)
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
