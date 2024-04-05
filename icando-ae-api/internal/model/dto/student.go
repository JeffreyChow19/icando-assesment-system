package dto

import (
	"github.com/google/uuid"
)

type GetAllStudentsFilter struct {
	Name               *string `form:"name"`
	InstitutionID      *string `form:"institutionId"`
	ClassID            *string `form:"classId"`
	Page               int     `form:"page"`
	Limit              int     `form:"limit"`
	IncludeInstitution bool
	IncludeClass       bool
	OrderBy            *string `form:"orderBy"`
	Asc                bool    `form:"asc"`
}

type GetStudentFilter struct {
	ID                 *string
	Nisn               *string
	Email              *string
	IncludeInstitution bool
	IncludeClass       bool
}

type CreateStudentDto struct {
	FirstName string     `json:"firstName" binding:"required"`
	LastName  string     `json:"lastName" binding:"required"`
	Nisn      string     `json:"nisn" binding:"required"`
	Email     string     `json:"email" binding:"required,email"`
	ClassID   *uuid.UUID `json:"classId"`
}

type UpdateStudentDto struct {
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	ClassID   *uuid.UUID `json:"classId"`
}

type UpdateStudentClassIdDto struct {
	StudentIDs []uuid.UUID `json:"studentIds" binding:"required"`
	ClassID    *uuid.UUID  `json:"classId" binding:"required"`
}
