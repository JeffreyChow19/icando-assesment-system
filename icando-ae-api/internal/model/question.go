package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"icando/utils"
)

type Question struct {
	Model
	Choices  *postgres.Jsonb `gorm:"type:jsonb"`
	AnswerID int             `gorm:"column:answer_id"`
	QuizID   uuid.UUID       `gorm:"column:quiz_id"`
}

type QuestionChoice struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (q *Question) GetQuestionChoices() (choices []QuestionChoice, err error) {
	err = utils.Unmarshal(q.Choices.RawMessage, &choices)
	return choices, err
}

func (q *Question) SetQuestionChoices(choices []QuestionChoice) error {
	jsonChoices, err := utils.Marshal(choices)
	if err != nil {
		return err
	}
	q.Choices = &postgres.Jsonb{RawMessage: jsonChoices}
	return nil
}
