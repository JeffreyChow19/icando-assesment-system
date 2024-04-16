package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"net/http"
	"time"
)

type StudentQuizService interface {
	UpdateStudentAnswer(studentQuiz *model.StudentQuiz, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError
}

type StudentQuizServiceImpl struct {
	studentQuizRepository repository.StudentQuizRepository
	questionRepository    repository.QuestionRepository
	quizRepository        repository.QuizRepository
}

func NewStudentQuizServiceImpl(studentQuizRepository repository.StudentQuizRepository, questionRepository repository.QuestionRepository, quizRepository repository.QuizRepository) *StudentQuizServiceImpl {
	return &StudentQuizServiceImpl{
		studentQuizRepository: studentQuizRepository,
		questionRepository:    questionRepository,
		quizRepository:        quizRepository,
	}
}

var ErrInvalidQuestion = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invalid question"),
}
var ErrStudentQuizNotStarted = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Student Quiz Not Started"),
}
var ErrStudentQuizSubmitted = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Student Quiz Submitted"),
}
var ErrInvalidQuizAttemptTime = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Invalid Quiz Attempt Time"),
}
var ErrUpdateStudentAnswer = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating student answer"),
}

func (s *StudentQuizServiceImpl) UpdateStudentAnswer(studentQuiz *model.StudentQuiz, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError {
	if studentQuiz.Status == enum.NOT_STARTED {
		return ErrStudentQuizNotStarted
	}

	if studentQuiz.Status == enum.SUBMITTED {
		return ErrStudentQuizSubmitted
	}

	quiz, errQuiz := s.quizRepository.GetQuiz(dto.GetQuizFilter{
		ID: studentQuiz.QuizID,
	})

	if errQuiz != nil {
		if errors.Is(errQuiz, gorm.ErrRecordNotFound) {
			return ErrQuizNotFound
		}
		return ErrUpdateStudentAnswer
	}

	currentTime := time.Now()
	if quiz.EndAt != nil && currentTime.After(*quiz.EndAt) {
		return ErrInvalidQuizAttemptTime
	}
	if quiz.StartAt != nil && currentTime.Before(*quiz.StartAt) {
		return ErrInvalidQuizAttemptTime
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
		StudentQuizID: studentQuiz.ID,
		QuestionID:    questionID,
		AnswerID:      studentAnswerDto.AnswerID,
	}

	err := s.studentQuizRepository.UpdateAnswer(studentAnswer)
	if err != nil {
		return ErrUpdateStudentAnswer
	}

	return nil
}
