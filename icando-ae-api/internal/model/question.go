package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
)

type Question struct {
	Model
	Text         string
	Choices      *postgres.Jsonb `gorm:"type:jsonb"`
	AnswerID     int             `gorm:"column:answer_id"`
	QuizID       uuid.UUID       `gorm:"column:quiz_id"`
	Competencies []Competency    `gorm:"many2many:question_competencies;"`
}

type QuestionChoice struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type QuestionCompetency struct {
	QuestionID   uuid.UUID
	CompetencyID uuid.UUID
}

func (q *Question) GetQuestionChoices() (choices []QuestionChoice, err error) {
	err = json.Unmarshal(q.Choices.RawMessage, &choices)
	return choices, err
}

func (q *Question) SetQuestionChoices(choices []dto.QuestionChoiceDto) error {
	jsonChoices, err := json.Marshal(choices)
	if err != nil {
		return err
	}
	q.Choices = &postgres.Jsonb{RawMessage: jsonChoices}
	return nil
}

func (qc QuestionChoice) ToDao() dao.QuestionChoiceDao {
	return dao.QuestionChoiceDao{
		ID:   qc.ID,
		Text: qc.Text,
	}
}

func (q Question) ToDao() (*dao.QuestionDao, error) {
	choices, err := q.GetQuestionChoices()
	if err != nil {
		return nil, err
	}

	// Convert choices to a slice of QuestionChoiceDao
	var choicesDao []dao.QuestionChoiceDao
	for _, choice := range choices {
		choicesDao = append(choicesDao, choice.ToDao())
	}

	// Convert competencies to a slice of CompetencyDao
	var competenciesDao []dao.CompetencyDao
	for _, competency := range q.Competencies {
		competencyDao := dao.CompetencyDao{
			ID:          competency.ID,
			Numbering:   &competency.Numbering,
			Name:        &competency.Name,
			Description: &competency.Description,
		}
		competenciesDao = append(competenciesDao, competencyDao)
	}

	daoQuiz := dao.QuestionDao{
		ID:           q.ID,
		Text:         q.Text,
		Choices:      choicesDao,
		AnswerID:     q.AnswerID,
		Competencies: competenciesDao,
	}

	return &daoQuiz, nil
}
