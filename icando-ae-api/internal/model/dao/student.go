package dao

import (
	"github.com/google/uuid"
	"time"
)

type StudentDao struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Nisn      string    `json:"nisn"`
	Email     string    `json:"email,omitempty"`
	ClassID   *uuid.UUID `json:"classId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// can add optional relations when needed
}
