package dto

import (
	"github.com/google/uuid"
	"time"
)

type GetQuizFilter struct {
	ID uuid.UUID
}

type UpdateQuizDto struct {
	ID           uuid.UUID  `json:"id" binding:"required"`
	Name         *string    `json:"name"`
	Subject      *string    `json:"subject"`
	PassingGrade float64    `json:"passing_grade"`
	Deadline     *time.Time `json:"deadline"`
}
