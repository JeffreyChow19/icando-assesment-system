package middleware

import "go.uber.org/fx"

var Module = fx.Module(
	"middleware",
	fx.Options(
		fx.Provide(NewAuthMiddleware),
	),
)
