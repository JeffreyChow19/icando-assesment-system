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

func (r *QuestionCompetencyRepository) Delete(questionID uuid.UUID) error {
	query := r.db.Model(&model.QuestionCompetency{})
	err := query.Where("question_id = ?", questionID).Delete(&model.QuestionCompetency{}).Error
	if err != nil {
		return err
	}
	return nil
}
