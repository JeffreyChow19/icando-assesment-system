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
			fx.Annotate(designer.NewStudentHandlerImpl, fx.As(new(designer.StudentHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewClassHandlerImpl, fx.As(new(designer.ClassHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewCompetencyHandlerImpl, fx.As(new(designer.CompetencyHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewQuizHandlerImpl, fx.As(new(designer.QuizHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewQuestionHandlerImpl, fx.As(new(designer.QuestionHandler))),
		),
		fx.Provide(
			fx.Annotate(designer.NewTeacherHandlerImpl, fx.As(new(designer.TeacherHandler))),
		),
	),
)
