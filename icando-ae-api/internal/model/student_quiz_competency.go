package model

import "github.com/google/uuid"

type StudentQuizCompetency struct {
	StudentQuizID uuid.UUID `gorm:"primarykey;column:student_quiz_id"`
	StudentID     uuid.UUID `gorm:"primarykey;column:student_id"`
	CompetencyID  uuid.UUID `gorm:"primaryKey;column:competency_id"`
	CorrectCount  int
	TotalCount    int
}
