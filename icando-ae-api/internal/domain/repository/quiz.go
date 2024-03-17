package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *lib.Database) QuizRepository {
	return QuizRepository{
		db: db.DB,
	}
}

func (r *QuizRepository) GetQuiz(filter dto.GetQuizFilter) (*model.Quiz, error) {
	query := r.db.Model(&model.Quiz{})

	if filter.ID != uuid.Nil {
		query.Where("id = ?", filter.ID)
	}

	var quiz model.Quiz
	err := query.First(&quiz).Error
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (r *QuizRepository) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	err := r.db.Create(&quiz).Error
	return quiz, err
}

func (r *QuizRepository) UpdateQuiz(quiz model.Quiz) error {
	return r.db.Save(&quiz).Error
}
