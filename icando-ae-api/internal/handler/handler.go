package handler

import "go.uber.org/fx"

var Module = fx.Module(
	"handler",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewHealthcheckHandlerImpl, fx.As(new(HealthcheckHandler))),
		),
		fx.Provide(
			fx.Annotate(NewStudentHandlerImpl, fx.As(new(StudentHandler))),
		),
	),
)
