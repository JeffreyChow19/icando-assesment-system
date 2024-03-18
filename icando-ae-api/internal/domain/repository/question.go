package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *lib.Database) QuestionRepository {
	return QuestionRepository{
		db: db.DB,
	}
}

func (r *QuestionRepository) GetQuestion(filter dto.GetQuestionFilter) (*model.Question, error) {
	query := r.db.Model(&model.Question{})

	if filter.ID != uuid.Nil {
		query = query.Where("id = ?", filter.ID)
	}

	if filter.QuizID != uuid.Nil {
		query = query.Where("quiz_id = ?", filter.QuizID)
	}

	var question model.Question
	err := query.First(&question).Error
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *QuestionRepository) CreateQuestion(question model.Question) (model.Question, error) {
	err := r.db.Set("gorm:association_autoupdate", false).Create(&question).Error
	return question, err
}

func (r *QuestionRepository) UpdateQuestion(question model.Question) error {
	return r.db.Save(&question).Error
}
