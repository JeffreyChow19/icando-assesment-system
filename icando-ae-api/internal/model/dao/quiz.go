package dao

import (
	"github.com/google/uuid"
	"time"
)

type QuizDao struct {
	ID           uuid.UUID  `json:"id"`
	Name         *string    `json:"name,omitempty"`
	Subject      *string    `json:"subject,omitempty"`
	PassingGrade float64    `json:"passing_grade,omitempty"`
	Deadline     *time.Time `json:"deadline,omitempty"`
}
