package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/lib"
)

type QuestionCompetencyRepository struct {
	db *gorm.DB
}

func NewQuestionCompetencyRepository(db *lib.Database) QuestionCompetencyRepository {
	return QuestionCompetencyRepository{
		db: db.DB,
	}
}

func (r *QuestionCompetencyRepository) GetAll(questionID uuid.UUID) ([]model.QuestionCompetency, error) {
	var questionCompetencies []model.QuestionCompetency

	err := r.db.Where("question_id = ?", questionID).Find(&questionCompetencies).Error
	if err != nil {
		return nil, err
	}

	return questionCompetencies, nil
}

func (r *QuestionCompetencyRepository) Delete(questionCompetencies []model.QuestionCompetency) error {
	var questionIDs []uuid.UUID
	var competencyIDs []uuid.UUID
	for _, qc := range questionCompetencies {
		questionIDs = append(questionIDs, qc.QuestionID)
		competencyIDs = append(competencyIDs, qc.CompetencyID)
	}

	if err := r.db.Exec("DELETE FROM question_competencies WHERE question_id IN (?) AND competency_id IN (?)", questionIDs, competencyIDs).Error; err != nil {
		return err
	}

	return nil
}
