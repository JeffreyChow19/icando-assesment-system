package repository

import (
	"gorm.io/gorm"
	"icando/internal/model"
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

func (r *QuizRepository) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	err := r.db.Create(&quiz).Error
	return quiz, err
}
