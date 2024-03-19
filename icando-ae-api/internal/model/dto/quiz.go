package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetQuizFilter struct {
	ID uuid.UUID
	WithCreator		bool
	WithUpdater 	bool
	WithQuestions	bool
}

type UpdateQuizDto struct {
	ID           uuid.UUID  `json:"id" binding:"required"`
	Name         *string    `json:"name"`
	Subject      *string    `json:"subject"`
	PassingGrade float64    `json:"passing_grade"`
	Deadline     *time.Time `json:"deadline"`
}
