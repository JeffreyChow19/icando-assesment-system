package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type StudentQuizService interface {
	UpdateStudentAnswer(studentQuizID uuid.UUID, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError
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

var ErrStudentQuizNotFound = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Student Quiz Not Found"),
}
var ErrUpdateStudentAnswer = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating student answer"),
}

func (s *StudentQuizServiceImpl) UpdateStudentAnswer(studentQuizID uuid.UUID, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError {
	_, errGetStudentQuiz := s.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{
		ID: studentQuizID,
	})
	if errGetStudentQuiz != nil {
		return ErrStudentQuizNotFound
	}

	_, errGetQuestion := s.questionRepository.GetQuestion(dto.GetQuestionFilter{
		ID: questionID,
	})
	if errGetQuestion != nil {
		return ErrQuestionNotFound
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
