package dao

import (
	"github.com/google/uuid"
	"icando/internal/model/base"
	"time"
)

type QuizDao struct {
	ID           uuid.UUID        `json:"id"`
	Name         *string          `json:"name"`
	Subject      base.StringArray `json:"subject" gorm:"type:text[]"`
	PassingGrade float64          `json:"passingGrade"`
	PublishedAt  *time.Time       `json:"publishedAt"`
	Duration     *int             `json:"duration"`
	StartAt      *time.Time       `json:"startAt"`
	EndAt        *time.Time       `json:"endAt"`
	Creator      *TeacherDao      `json:"creator,omitempty"`
	Updater      *TeacherDao      `json:"updater,omitempty"`
	Questions    []QuestionDao    `json:"questions"`
	Classes      []ClassDao       `json:"classes,omitempty"`
}

type ParentQuizDao struct {
	ID              uuid.UUID        `json:"id"`
	Name            *string          `json:"name"`
	Subject         base.StringArray `json:"subject" gorm:"type:text[]"`
	PassingGrade    float64          `json:"passingGrade"`
	LastPublishedAt *time.Time       `json:"lastPublishedAt"`
	CreatedBy       string           `json:"createdBy"`
	UpdatedBy       string           `json:"updatedBy"`
}
