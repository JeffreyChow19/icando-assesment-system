package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"math"
	"strings"
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

func (r *QuizRepository) GetAllQuiz(filter dto.GetAllQuizzesFilter) ([]model.Quiz, *dao.MetaDao, error) {
	query := r.db.Table("quizzes").Where("parent_quiz IS NULL")
	if filter.Query != nil {
		query.Where("LOWER(name) LIKE ?", strings.ToLower(fmt.Sprintf("%%%s%%", *filter.Query)))
	}
	if filter.Subject != nil {
		query.Where("LOWER(subject) LIKE ?", strings.ToLower(fmt.Sprintf("%%%s%%", *filter.Subject)))
	}

	var totalItem int64
	err := query.Session(&gorm.Session{}).Count(&totalItem).Error
	if err != nil {
		return nil, nil, err
	}

	meta := dao.MetaDao{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(float64(totalItem) / float64(filter.Limit))),
	}
	Paginate(query, filter.Page, filter.Limit)

	var quizzes []model.Quiz
	err = query.Session(&gorm.Session{}).Find(&quizzes).Error

	return quizzes, &meta, err
}
