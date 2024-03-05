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
	fx.Provide(NewStudentRoute),
)

type Route interface {
	Setup(engine *gin.RouterGroup)
}

func NewRoutes(
	//fileRoute *FileRoute,
	healthcheckRoute *HealthcheckRoute,
	classRoute *StudentRoute,
) *Routes {
	publicRoutes := []Route{healthcheckRoute}
	designerRoute := []Route{classRoute}
	return &Routes{
		public:   publicRoutes,
		designer: designerRoute,
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
