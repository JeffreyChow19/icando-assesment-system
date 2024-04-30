package dto

import "github.com/google/uuid"

type GetStudentCompetencyFilter struct {
	StudentID     *uuid.UUID
	StudentQuizID *uuid.UUID
}
