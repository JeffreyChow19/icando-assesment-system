package model

import (
	"github.com/google/uuid"
)

type Student struct {
	ID            uuid.UUID `gorm:"primarykey"`
	FirstName     string
	LastName      string
	Nisn          string
	Email         string
	Password      string
	InstitutionID uuid.UUID
	ClassID       uuid.UUID
	Institution   *Institution
	Class         *Class
}
