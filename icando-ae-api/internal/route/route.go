package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Routes struct {
	student  []Route
	teacher  []Route
	designer []Route
	public   []Route
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
	//fx.Provide(NewFileRoute),
	fx.Provide(NewHealthcheckRoute),
)

type Route interface {
	Setup(engine *gin.RouterGroup)
}

func NewRoutes(
	//fileRoute *FileRoute,
	healthcheckRoute *HealthcheckRoute,
) *Routes {
	publicRoutes := []Route{healthcheckRoute}
	return &Routes{
		public: publicRoutes,
	}
}

func (r Routes) Setup(engine *gin.Engine) {
	public := engine.Group("")
	student := engine.Group("/student")
	teacher := engine.Group("/teacher")
	designer := engine.Group("/designer")
	for _, route := range r.public {
		route.Setup(public)
	}
	for _, route := range r.student {
		route.Setup(student)
	}
	for _, route := range r.teacher {
		route.Setup(teacher)
	}
	for _, route := range r.designer {
		route.Setup(designer)
	}
}
