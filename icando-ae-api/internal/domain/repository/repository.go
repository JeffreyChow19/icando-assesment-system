package repository

import "go.uber.org/fx"

var Module = fx.Module(
	"repository",
	fx.Options(
		fx.Provide(NewClassRepository),
	),
)
