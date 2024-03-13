package service

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewCompetencyServiceImpl, fx.As(new(CompetencyService))),
		),
		fx.Provide(
			fx.Annotate(NewStudentServiceImpl, fx.As(new(StudentService))),
		),
		fx.Provide(fx.Annotate(NewAuthServiceImpl, fx.As(new(AuthService)))),
		fx.Provide(fx.Annotate(NewTeacherServiceImpl, fx.As(new(TeacherService)))),
		fx.Provide(fx.Annotate(NewClassServiceImpl, fx.As(new(ClassService)))),
	),
)
