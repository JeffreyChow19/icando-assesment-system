package dao

import (
	"github.com/google/uuid"
)

type TeacherDao struct {
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         *string   `json:"email"`
	InstitutionID uuid.UUID `json:"institutionId"`
}
