package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type StudentQuizRepository struct {
	db             *gorm.DB
	quizRepository *QuizRepository
}

func NewStudentQuizRepository(db *lib.Database, quizRepository *QuizRepository) StudentQuizRepository {
	return StudentQuizRepository{
		db:             db.DB,
		quizRepository: quizRepository,
	}
}

func (r *StudentQuizRepository) GetStudentQuiz(filter dto.GetStudentQuizFilter) (*model.StudentQuiz, error) {
	query := r.db.Session(&gorm.Session{})

	if filter.WithAnswers {
		query.Preload("StudentAnswers")
	}

	query = query.Where("id = ?", filter.ID.String())

	var studentQuiz model.StudentQuiz

	if err := query.First(&studentQuiz).Error; err != nil {
		return nil, err
	}

	if filter.WithQuizOverview || filter.WithQuizQuestions {
		quiz, err := r.quizRepository.GetQuiz(dto.GetQuizFilter{
			ID:            studentQuiz.QuizID,
			WithQuestions: filter.WithQuizQuestions,
		})

		if err != nil {
			return nil, err
		}

		studentQuiz.Quiz = quiz
	}

	return &studentQuiz, nil
}

func (r *StudentQuizRepository) CreateStudentQuiz(studentQuiz model.StudentQuiz) (model.StudentQuiz, error) {
	err := r.db.Create(&studentQuiz).Error

	return studentQuiz, err
}

func (r *StudentQuizRepository) UpdateStudentQuiz(studentQuiz model.StudentQuiz) error {
	return r.db.Save(&studentQuiz).Error
}

func (r *StudentQuizRepository) UpdateAnswer(answer model.StudentAnswer) error {
	return r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&answer).Error
}
