package dto

import (
	"github.com/google/uuid"
	"icando/internal/model/enum"
)

type GetTeacherFilter struct {
	ID            *uuid.UUID
	Email         *string
	Role          *enum.TeacherRole
	InstitutionID *uuid.UUID
	WithClasses   *bool
}
