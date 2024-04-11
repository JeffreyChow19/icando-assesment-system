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
	PassingGrade float64    `json:"passingGrade"`
	Deadline     *time.Time `json:"deadline"`
}

type PublishQuizDto struct {
	QuizID          uuid.UUID   `json:"quizId" binding:"required"`
	QuizDuration    int         `json:"quizDuration" binding:"required" validate:"gt=0"`
	StartDate       time.Time   `json:"startDate" binding:"required"`
	EndDate         time.Time   `json:"endDate" binding:"required"`
	AssignedClasses []uuid.UUID `json:"assignedClasses" binding:"required" validate:"min=1"`
}
