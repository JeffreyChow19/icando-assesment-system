package dto

import "github.com/google/uuid"

type GetStudentQuizFilter struct {
	ID                uuid.UUID
	WithQuizOverview  bool
	WithQuizQuestions bool
	WithAnswers       bool
	WithStudent       bool
}

type UpdateStudentAnswerDto struct {
	AnswerID int `json:"answer_id"`
}

type StudentQuizCompetencyCorrectTotalDto struct {
	TotalCount   int
	CorrectCount int
}
