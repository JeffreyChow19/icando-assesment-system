package dao

import "github.com/google/uuid"

type ClassDao struct {
	ID           uuid.UUID       `json:"id"`
	Name         string          `json:"name"`
	Grade        string          `json:"grade"`
	InstituionID uuid.UUID       `json:"instituionId"`
	Teachers     []TeacherDao    `json:"teachers,omitempty"`
	Institution  *InstitutionDao `json:"institution,omitempty"`
	Students     []StudentDao    `json:"students,omitempty"`
}
