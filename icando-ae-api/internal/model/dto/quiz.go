package dto

import (
	"github.com/google/uuid"
	"time"
)

type GetQuizFilter struct {
	ID uuid.UUID
}

type GetAllQuizzesFilter struct {
	InstitutionID *string `form:"institutionId"`
	Query         *string `form:"q"`
	Subject       *string `form:"subject"`
	Page          int     `form:"page"`
	Limit         int     `form:"limit"`
}

type UpdateQuizDto struct {
	ID           uuid.UUID  `json:"id" binding:"required"`
	Name         *string    `json:"name"`
	Subject      *string    `json:"subject"`
	PassingGrade float64    `json:"passing_grade"`
	Deadline     *time.Time `json:"deadline"`
}
