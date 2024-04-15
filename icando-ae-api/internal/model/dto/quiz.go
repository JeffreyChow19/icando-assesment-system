package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetQuizFilter struct {
	ID            uuid.UUID
	WithCreator   bool
	WithUpdater   bool
	WithQuestions bool
	WithClasses   bool
}

type GetAllQuizzesFilter struct {
	InstitutionID *string  `form:"institutionId"`
	Query         *string  `form:"q"`
	Subject       []string `form:"subject"`
	Page          int      `form:"page"`
	Limit         int      `form:"limit"`
}

type UpdateQuizDto struct {
	ID           uuid.UUID  `json:"id" binding:"required"`
	Name         *string    `json:"name"`
	Subject      []string   `json:"subject"`
	PassingGrade float64    `json:"passingGrade"`
	StartAt	     *time.Time `json:"startAt"`
	EndAt		     *time.Time `json:"endAt"`
}
