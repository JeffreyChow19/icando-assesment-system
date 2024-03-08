package dto

import "github.com/google/uuid"

type GetTeacherFilter struct {
	ID    *uuid.UUID
	Email *string
}
