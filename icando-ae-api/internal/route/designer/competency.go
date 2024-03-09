package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
)

type CompetencyRoute struct {
	competencyHandler designer.CompetencyHandler
}

func (r CompetencyRoute) Setup(engine *gin.RouterGroup) {
	group := engine.Group("/competency")
	group.GET("", r.competencyHandler.GetAllCompetencies)
	group.POST("", r.competencyHandler.CreateCompetency)
	group.PATCH("", r.competencyHandler.UpdateCompetency)
	group.DELETE("/:id", r.competencyHandler.DeleteCompetency)
}

func NewCompetencyRoute(handler designer.CompetencyHandler) *CompetencyRoute {
	return &CompetencyRoute{
		competencyHandler: handler,
	}
}
