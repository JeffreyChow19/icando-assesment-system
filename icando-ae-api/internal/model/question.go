package model

import "github.com/google/uuid"
import "encoding/json"

type Question struct {
	Model
	Numbering          string
	Choices            *json.RawMessage `gorm:"type:jsonb"`
	CorrectChoiceIndex int
	QuizID             uuid.UUID
}
