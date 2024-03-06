package service

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewCompetencyServiceImpl, fx.As(new(CompetencyService))),
		),
	),
)
