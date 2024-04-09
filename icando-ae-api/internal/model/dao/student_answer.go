package dao

import (
	"github.com/google/uuid"
	"time"
)

type StudentAnswerCompetencyDao struct {
	CompetencyID uuid.UUID `json:"competency_id"`
	IsPassed     bool      `json:"is_passed"`
}

type StudentAnswerDao struct {
	ID            uuid.UUID                    `json:"id"`
	CreatedAt     time.Time                    `json:"createdAt"`
	UpdatedAt     time.Time                    `json:"updatedAt"`
	QuestionID    uuid.UUID                    `json:"questionId"`
	Question      *QuestionDao                 `json:"question"`
	StudentQuizID uuid.UUID                    `json:"studentQuizId"`
	AnswerID      int                          `json:"answerId"`
	IsCorrect     *bool                        `json:"isCorrect"`
	Competencies  []StudentAnswerCompetencyDao `json:"competencies"`
}
