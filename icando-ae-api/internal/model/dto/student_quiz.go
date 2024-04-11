package dto

import "github.com/google/uuid"

type GetStudentQuizFilter struct {
	ID                uuid.UUID
	WithQuizOverview  bool
	WithQuizQuestions bool
	WithAnswers       bool
}
