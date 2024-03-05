package model

import "github.com/google/uuid"

type Competency struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Numbering   string
	Name        string
	Description string
}
