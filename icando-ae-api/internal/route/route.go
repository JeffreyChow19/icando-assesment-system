package route

import (
	"icando/internal/middleware"
	"icando/internal/model/enum"
	"icando/internal/route/designer"
	"icando/internal/route/student"
	"icando/internal/route/teacher"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Routes struct {
	student        []Route
	teacher        []Route
	designer       []Route
	public         []Route
	auth           []Route
	authMiddleware *middleware.AuthMiddleware
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
	fx.Provide(designer.NewQuizRoute),
	fx.Provide(designer.NewQuestionRoute),
	fx.Provide(designer.NewTeacherRoute),
	fx.Provide(student.NewQuizRoute),
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
	quizRoute *designer.QuizRoute,
	questionRoute *designer.QuestionRoute,
	teacherRoute *designer.TeacherRoute,
	studentQuizRoute *student.QuizRoute,
	authMiddleware *middleware.AuthMiddleware,
) *Routes {
	publicRoutes := []Route{healthcheckRoute}
	designerRoutes := []Route{studentRoute, designerClass, competencyRoute, quizRoute, questionRoute, teacherRoute}
	teacherRoutes := []Route{}
	studentRoutes := []Route{studentQuizRoute}
	authRoutes := []Route{teacherAuth, designerAuth, studentAuth}

	return &Routes{
		public:         publicRoutes,
		designer:       designerRoutes,
		teacher:        teacherRoutes,
		student:        studentRoutes,
		authMiddleware: authMiddleware,
		auth:           authRoutes,
	}
}

func (r Routes) Setup(engine *gin.Engine) {
	public := engine.Group("")

	studentGroup := engine.Group("/student")
	studentGroup.Use(r.authMiddleware.Handler(enum.ROLE_STUDENT))

	teacherGroup := engine.Group("/teacher")
	teacherGroup.Use(r.authMiddleware.Handler(enum.ROLE_TEACHER))

	designerGroup := engine.Group("/designer")
	designerGroup.Use(r.authMiddleware.Handler(enum.ROLE_LEARNING_DESIGNER))

	authGroup := engine.Group("/auth")

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

	for _, route := range r.auth {
		route.Setup(authGroup)
	}
}
