package handler

import (
	"icando/internal/handler/designer"
	"icando/internal/handler/student"
	"icando/internal/handler/teacher"
	"go.uber.org/fx"
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
		fx.Provide(
			fx.Annotate(student.NewQuizHandlerImpl, fx.As(new(student.QuizHandler))),
		),
		fx.Provide(
			fx.Annotate(teacher.NewQuizHandlerImpl, fx.As(new(teacher.QuizHandler))),
		),
		fx.Provide(
			fx.Annotate(teacher.NewAnalyticsHandlerImpl, fx.As(new(teacher.AnalyticsHandler))),
		),
	),
)
