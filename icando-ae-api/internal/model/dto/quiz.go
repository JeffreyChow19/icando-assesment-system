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

type GetQuizVersionFilter struct {
	ID    uuid.UUID
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
type GetAllQuizzesFilter struct {
	TeacherID     *uuid.UUID
	InstitutionID *string  `form:"institutionId"`
	Query         *string  `form:"name"`
	Subject       []string `form:"subject"`
	Page          int      `form:"page"`
	Limit         int      `form:"limit"`
}

type UpdateQuizDto struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	Name         *string   `json:"name"`
	Subject      []string  `json:"subject"`
	PassingGrade float64   `json:"passingGrade"`
}

type PublishQuizDto struct {
	QuizID          uuid.UUID   `json:"quizId" binding:"required"`
	QuizDuration    int         `json:"quizDuration" binding:"required" validate:"gt=0"`
	StartAt         time.Time   `json:"startAt" binding:"required"`
	EndAt           time.Time   `json:"endAt" binding:"required"`
	AssignedClasses []uuid.UUID `json:"assignedClasses" binding:"required" validate:"min=1"`
}
