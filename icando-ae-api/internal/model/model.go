package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamptz"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamptz"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}
