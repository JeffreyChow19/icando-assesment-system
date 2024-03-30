package dao

import (
	"github.com/google/uuid"
	"time"
)

type QuizDao struct {
	ID           uuid.UUID     `json:"id"`
	Name         *string       `json:"name"`
	Subject      *string       `json:"subject"`
	PassingGrade float64       `json:"passingGrade"`
	PublishedAt  *time.Time    `json:"publishedAt"`
	Deadline     *time.Time    `json:"deadline"`
	Creator      *TeacherDao   `json:"creator,omitempty"`
	Updater      *TeacherDao   `json:"updater,omitempty"`
	Questions    []QuestionDao `json:"questions"`
}

type ParentQuizDao struct {
	ID              uuid.UUID  `json:"id"`
	Name            *string    `json:"name"`
	Subject         *string    `json:"subject"`
	PassingGrade    float64    `json:"passingGrade"`
	LastPublishedAt *time.Time `json:"lastPublishedAt"`
	CreatedBy       string     `json:"createdBy"`
}
