package dto

import "github.com/google/uuid"

type ClassDto struct {
	Name          string      `json:"name" binding:"required"`
	Grade         string      `json:"grade" binding:"required,numeric"`
	InstitutionID uuid.UUID   `json:"institutionId" binding:"required"`
	TeacherIDs    []uuid.UUID `json:"teacherIds" binding:"required"`
}

type CreateUpdateClassPayload struct {
	Name          string      `json:"name" binding:"required"`
	Grade         string      `json:"grade" binding:"required,numeric"`
	TeacherIDs    []uuid.UUID `json:"teacherIds" binding:"required"`
}

type GetAllClassFilter struct {
	InstitutionID *uuid.UUID `json:"institutionId"`
	TeacherID     *uuid.UUID `json:"teacherId"`
	SortBy        *string    `json:"sortBy"`
	Desc          bool       `json:"desc"`
}

type GetClassFilter struct {
	WithTeacherRelation     bool
	WithInstitutionRelation bool
	WithStudentRelation     bool
}

type AssignStudentsRequest struct {
	StudentIDs []uuid.UUID `json:"studentIds" binding:"required"`
}
