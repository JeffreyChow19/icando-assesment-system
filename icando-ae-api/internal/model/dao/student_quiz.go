package dao

import (
	"github.com/google/uuid"
	"icando/internal/model/enum"
	"time"
)

type StudentQuizDao struct {
	ID             uuid.UUID          `json:"id"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	TotalScore     int                `json:"totalScore"`
	StartedAt      *time.Time         `json:"startedAt"`
	CompletedAt    *time.Time         `json:"completedAt"`
	Status         enum.QuizStatus    `json:"status"`
	QuizID         uuid.UUID          `json:"quiz_id"`
	Quiz           *QuizDao           `json:"quiz,omitempty"`
	StudentAnswers []StudentAnswerDao `json:"studentAnswers"`
}
