package handler

import "go.uber.org/fx"

var Module = fx.Module("handler", fx.Options(fx.Provide(NewEmailHandler), fx.Provide(NewScoreHandler)))
