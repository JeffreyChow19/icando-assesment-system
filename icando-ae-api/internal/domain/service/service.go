package service

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewStudentServiceImpl, fx.As(new(StudentService))),
		),
		fx.Provide(fx.Annotate(NewAuthServiceImpl, fx.As(new(AuthService)))),
		fx.Provide(fx.Annotate(NewLearningDesignerServiceImpl, fx.As(new(LearningDesignerService)))),
	),
)
