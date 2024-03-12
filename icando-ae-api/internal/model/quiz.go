package model

import (
	"github.com/google/uuid"
	"time"
)

type Quiz struct {
	Model
	Name         string
	Subject      string
	PassingGrade float64
	ParentQuiz   uuid.UUID
	CreatedBy    uuid.UUID `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID `gorm:"type:uuid"`
	PublishedAt  time.Time `gorm:"type:timestamptz"`
}
