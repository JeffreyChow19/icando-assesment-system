package dto

import "github.com/google/uuid"

type GetAllStudentsFilter struct {
	Name               *string
	InstitutionID      *uuid.UUID
	ClassID            *uuid.UUID
	Page               int
	Limit              int
	IncludeInstitution bool
	IncludeClass       bool
}

type CreateStudentDto struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Nisn      string    `json:"nisn" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	ClassID   uuid.UUID `json:"classId" binding:"required"`
}
