package route

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
)

type CompetencyRoute struct {
	competencyHandler handler.CompetencyHandler
}

func (r CompetencyRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/competency")
	group.GET("", r.competencyHandler.GetAllCompetencies)
	group.POST("", r.competencyHandler.CreateCompetency)
	group.PATCH("", r.competencyHandler.UpdateCompetency)
	group.DELETE("/:id", r.competencyHandler.DeleteCompetency)
}

func NewCompetencyRoute(handler handler.CompetencyHandler) *CompetencyRoute {
	return &CompetencyRoute{
		competencyHandler: handler,
	}
}
