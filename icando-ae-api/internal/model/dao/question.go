package dao

import (
	"github.com/google/uuid"
)

type QuestionDao struct {
	ID           uuid.UUID           `json:"id"`
	Text         string              `json:"text"`
	Choices      []QuestionChoiceDao `json:"choices"`
	AnswerID     *int                `json:"answerId,omitempty"`
	Competencies []CompetencyDao     `json:"competencies"`
	Order        int                 `json:"order"`
}

type QuestionChoiceDao struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
