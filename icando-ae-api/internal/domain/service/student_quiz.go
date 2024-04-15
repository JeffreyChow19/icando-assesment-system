package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type StudentQuizService interface {
	UpdateStudentAnswer(studentID uuid.UUID, studentQuizID uuid.UUID, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError
}

type StudentQuizServiceImpl struct {
	studentQuizRepository repository.StudentQuizRepository
	questionRepository    repository.QuestionRepository
}

func NewStudentQuizServiceImpl(studentQuizRepository repository.StudentQuizRepository, questionRepository repository.QuestionRepository) *StudentQuizServiceImpl {
	return &StudentQuizServiceImpl{
		studentQuizRepository: studentQuizRepository,
		questionRepository:    questionRepository,
	}
}

var ErrInvalidUser = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invalid user"),
}
var ErrInvalidQuestion = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invalid question"),
}
var ErrStudentQuizNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Student Quiz Not Found"),
}
var ErrUpdateStudentAnswer = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating student answer"),
}

func (s *StudentQuizServiceImpl) UpdateStudentAnswer(studentID uuid.UUID, studentQuizID uuid.UUID, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError {
	studentQuiz, errGetStudentQuiz := s.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{
		ID: studentQuizID,
	})
	if errGetStudentQuiz != nil {
		if errors.Is(errGetStudentQuiz, gorm.ErrRecordNotFound) {
			return ErrStudentQuizNotFound
		}
		return ErrUpdateStudentAnswer
	}

	if studentQuiz.StudentID != studentID {
		return ErrInvalidUser
	}

	question, errGetQuestion := s.questionRepository.GetQuestion(dto.GetQuestionFilter{
		ID: questionID,
	})
	if errGetQuestion != nil {
		if errors.Is(errGetQuestion, gorm.ErrRecordNotFound) {
			return ErrQuestionNotFound
		}
		return ErrUpdateStudentAnswer
	}

	if question.QuizID != studentQuiz.QuizID {
		return ErrInvalidQuestion
	}

	studentAnswer := model.StudentAnswer{
		StudentQuizID: studentQuizID,
		QuestionID:    questionID,
		AnswerID:      studentAnswerDto.AnswerID,
	}

	err := s.studentQuizRepository.UpdateAnswer(studentAnswer)
	if err != nil {
		return ErrUpdateStudentAnswer
	}

	return nil
}
