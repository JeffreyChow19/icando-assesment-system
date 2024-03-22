package repository

import "go.uber.org/fx"

var Module = fx.Module(
	"repository",
	fx.Options(
		fx.Provide(NewInstutionRepository),
		fx.Provide(NewStudentRepository),
		fx.Provide(NewTeacherRepository),
		fx.Provide(NewClassRepository),
		fx.Provide(NewCompetencyRepository),
		fx.Provide(NewQuizRepository),
		fx.Provide(NewQuestionRepository),
		fx.Provide(NewQuestionCompetencyRepository),
	),
)
