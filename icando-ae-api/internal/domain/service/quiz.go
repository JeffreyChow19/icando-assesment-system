package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/utils/httperror"
	"net/http"
)

type QuizService interface {
	CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
}

type QuizServiceImpl struct {
	quizRepository repository.QuizRepository
}

func NewQuizServiceImpl(quizRepository repository.QuizRepository) *QuizServiceImpl {
	return &QuizServiceImpl{
		quizRepository: quizRepository,
	}
}

var ErrCreateQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating quiz"),
}

func (s *QuizServiceImpl) CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError) {
	quiz := model.Quiz{
		CreatedBy: id,
	}
	quiz, err := s.quizRepository.CreateQuiz(quiz)

	if err != nil {
		return nil, ErrCreateQuiz
	}

	quizDao := quiz.ToDao()

	return &quizDao, nil
}
