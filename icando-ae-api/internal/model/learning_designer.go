package model

import (
	"github.com/google/uuid"
)

type LearningDesiner struct {
	ID        uuid.UUID `gorm:"primarykey"`
	FirstName 	string
	LastName 	string
	Email 	string
	Password 	string
	InstitutionID 	uuid.UUID
	Institution 	*Institution
}