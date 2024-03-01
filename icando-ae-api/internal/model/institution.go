package model

import (
	"github.com/google/uuid"
)

type Institution struct {
	ID   uuid.UUID `gorm:"primarykey"`
	Name string
	Nis  string
	Slug string
}
