package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type QuizService interface {
	CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	GetQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	UpdateQuiz(userID uuid.UUID, quizDto dto.UpdateQuizDto) (*dao.QuizDao, *httperror.HttpError)
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

var ErrQuizNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Quiz Not Found"),
}

var ErrGetQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when retrieving quiz"),
}

var ErrUpdateQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating quiz"),
}

func (s *QuizServiceImpl)	GetQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError) {
	quiz, err := s.quizRepository.GetQuiz(dto.GetQuizFilter{ID: id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrGetQuiz
	}

	quizDao := quiz.ToDao()
	return &quizDao, nil
}

func (s *QuizServiceImpl) UpdateQuiz(userID uuid.UUID, quizDto dto.UpdateQuizDto) (*dao.QuizDao, *httperror.HttpError) {
	quiz, err := s.quizRepository.GetQuiz(dto.GetQuizFilter{ID: quizDto.ID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrUpdateQuiz
	}

	quiz.UpdatedBy = &userID

	if quizDto.Name != nil {
		quiz.Name = quizDto.Name
	}
	if quizDto.Subject != nil {
		quiz.Subject = quizDto.Subject
	}
	if quizDto.PassingGrade != 0 {
		quiz.PassingGrade = quizDto.PassingGrade
	}
	if quizDto.Deadline != nil {
		quiz.Deadline = quizDto.Deadline
	}

	err = s.quizRepository.UpdateQuiz(*quiz)
	if err != nil {
		return nil, ErrUpdateQuiz
	}

	quizDao := quiz.ToDao()

	return &quizDao, nil
}
