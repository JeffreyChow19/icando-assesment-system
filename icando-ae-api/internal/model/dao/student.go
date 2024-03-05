package dao

import "github.com/google/uuid"

type StudentDao struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Nisn      string    `json:"nisn"`
	Email     *string   `json:"email,omitempty"`
	ClassID   uuid.UUID `json:"classId"`
	// can add optional relations when needed
}
