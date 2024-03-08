package handler

import (
	"go.uber.org/fx"
	"icando/internal/handler/designer"
)

var Module = fx.Module(
	"handler",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewHealthcheckHandlerImpl, fx.As(new(HealthcheckHandler))),
		),
		fx.Provide(
			fx.Annotate(NewAuthHandlerImpl, fx.As(new(AuthHandler))),
		),
		fx.Provide(
			fx.Annotate(NewStudentHandlerImpl, fx.As(new(StudentHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewClassHandlerImpl, fx.As(new(designer.ClassHandler))),
		),
		fx.Provide(
			fx.Annotate(NewCompetencyHandlerImpl, fx.As(new(CompetencyHandler))),
		),
	),
)
