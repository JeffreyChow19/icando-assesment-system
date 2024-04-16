package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"icando/internal/model/dao"
)

type StudentAnswer struct {
	Model
	QuestionID    uuid.UUID `gorm:"column:question_id"`
	Question      *Question
	StudentQuizID uuid.UUID `gorm:"column:student_quiz_id"`
	StudentQuiz   StudentQuiz
	AnswerID      int `gorm:"column:answer_id"`
	IsCorrect     *bool
	Competencies  *postgres.Jsonb `gorm:"type:jsonb"`
}

type StudentAnswerCompetency struct {
	CompetencyID uuid.UUID
	IsPassed     bool
}

func (sac StudentAnswerCompetency) ToDao() dao.StudentAnswerCompetencyDao {
	return dao.StudentAnswerCompetencyDao{
		CompetencyID: sac.CompetencyID,
		IsPassed:     sac.IsPassed,
	}
}

func (a *StudentAnswer) GetCompetencies() (competencies []StudentAnswerCompetency, err error) {
	err = json.Unmarshal(a.Competencies.RawMessage, &competencies)
	return competencies, err
}

func (a *StudentAnswer) SetCompetencies(competencies []StudentAnswerCompetency) error {
	jsonChoices, err := json.Marshal(competencies)
	if err != nil {
		return err
	}
	a.Competencies = &postgres.Jsonb{RawMessage: jsonChoices}
	return nil
}

func (a *StudentAnswer) ToDao(withQuestionAnswer bool) (*dao.StudentAnswerDao, error) {
	competencies, err := a.GetCompetencies()
	if err != nil {
		return nil, err
	}

	var competenciesDao []dao.StudentAnswerCompetencyDao
	for _, competency := range competencies {
		competenciesDao = append(competenciesDao, competency.ToDao())
	}

	daoAnswer := dao.StudentAnswerDao{
		ID:            a.ID,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
		QuestionID:    a.QuestionID,
		StudentQuizID: a.StudentQuizID,
		AnswerID:      a.AnswerID,
		IsCorrect:     a.IsCorrect,
		Competencies:  competenciesDao,
	}

	if a.Question != nil {
		questionDao, err := a.Question.ToDao(withQuestionAnswer)

		if err == nil {
			daoAnswer.Question = questionDao
		}
	}

	return &daoAnswer, nil
}
