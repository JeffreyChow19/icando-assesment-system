package dao

import "github.com/google/uuid"

type ClassDao struct {
	ID           uuid.UUID       `json:"id"`
	Name         string          `json:"name"`
	Grade        string          `json:"grade"`
	InstitutionID uuid.UUID      `json:"institutionId"`
	Teachers     []TeacherDao    `json:"teachers,omitempty"`
	Institution  *InstitutionDao `json:"institution,omitempty"`
	Students     []StudentDao    `json:"students,omitempty"`
}
