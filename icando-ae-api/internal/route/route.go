package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"icando/internal/route/designer"
	"icando/internal/route/student"
	"icando/internal/route/teacher"
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
	fx.Provide(designer.NewStudentRoute),
	fx.Provide(designer.NewAuthRoute),
	fx.Provide(teacher.NewAuthRoute),
	fx.Provide(student.NewAuthRoute),
	fx.Provide(designer.NewClassRoute),
	fx.Provide(designer.NewCompetencyRoute),
)

type Route interface {
	Setup(engine *gin.RouterGroup)
}

func NewRoutes(
	//fileRoute *FileRoute,
	healthcheckRoute *HealthcheckRoute,
	designerAuth *designer.AuthRoute,
	designerClass *designer.ClassRoute,
	teacherAuth *teacher.AuthRoute,
	studentAuth *student.AuthRoute,
	studentRoute *designer.StudentRoute,
	competencyRoute *designer.CompetencyRoute,
) *Routes {
	publicRoutes := []Route{healthcheckRoute}
	designerRoutes := []Route{studentRoute, designerAuth, designerClass, competencyRoute}
	teacherRoutes := []Route{teacherAuth}
	studentRoutes := []Route{studentAuth}

	return &Routes{
		public:   publicRoutes,
		designer: designerRoutes,
		teacher:  teacherRoutes,
		student:  studentRoutes,
	}
}

func (r Routes) Setup(engine *gin.Engine) {
	public := engine.Group("")

	studentGroup := engine.Group("/student")
	teacherGroup := engine.Group("/teacher")
	designerGroup := engine.Group("/designer")

	for _, route := range r.public {
		route.Setup(public)
	}

	for _, route := range r.student {
		route.Setup(studentGroup)
	}

	for _, route := range r.teacher {
		route.Setup(teacherGroup)
	}

	for _, route := range r.designer {
		route.Setup(designerGroup)
	}
}
