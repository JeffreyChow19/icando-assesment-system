package dao

import "github.com/google/uuid"

type CompetencyDao struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Numbering   *string   `json:"numbering,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
}
