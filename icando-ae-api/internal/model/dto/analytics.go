package dto

import "github.com/google/uuid"

type GetQuizPerformanceFilter struct {
	StudentID *uuid.UUID
	TeacherID *uuid.UUID
	QuizID    *uuid.UUID
}
