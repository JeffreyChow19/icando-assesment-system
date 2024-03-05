package service

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewStudentServiceImpl, fx.As(new(StudentService))),
		),
	),
)
