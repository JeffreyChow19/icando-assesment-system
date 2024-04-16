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
	CalculateScore(id uuid.UUID) error
}

type StudentQuizServiceImpl struct {
	studentQuizRepository repository.StudentQuizRepository
	questionRepository    repository.QuestionRepository
	quizRepository        repository.QuizRepository
	db                    *gorm.DB
}

func NewStudentQuizServiceImpl(
	studentQuizRepository repository.StudentQuizRepository,
	questionRepository repository.QuestionRepository,
	quizRepository repository.QuizRepository,
	db *gorm.DB) *StudentQuizServiceImpl {
	return &StudentQuizServiceImpl{
		studentQuizRepository: studentQuizRepository,
		questionRepository:    questionRepository,
		quizRepository:        quizRepository,
		db:                    db,
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

func (s *StudentQuizServiceImpl) CalculateScore(id uuid.UUID) error {
	studentQuiz, err := s.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{
		WithQuizQuestions: true,
		WithAnswers:       true,
		ID:                id,
	})

	if err != nil {
		return err
	}

	if studentQuiz.Status == enum.NOT_STARTED {
		return nil // we mark not started quiz as separated
	}

	answerMap := make(map[string]model.StudentAnswer)

	for _, answer := range studentQuiz.StudentAnswers {
		answerMap[answer.QuestionID.String()] = answer
	}

	questionCount := len(studentQuiz.Quiz.Questions)
	correctCount := 0

	for _, question := range studentQuiz.Quiz.Questions {
		answer, ok := answerMap[question.ID.String()]

		if ok {
			isCorrect := answer.AnswerID == question.AnswerID
			answer.IsCorrect = &isCorrect

			answerCompetencies := make([]model.StudentAnswerCompetency, 0)

			for _, competency := range question.Competencies {
				answerCompetencies = append(answerCompetencies, model.StudentAnswerCompetency{
					CompetencyID: competency.ID,
					IsPassed:     isCorrect,
				})
			}

			if err := answer.SetCompetencies(answerCompetencies); err != nil {
				return err
			}

			answerMap[question.ID.String()] = answer
			correctCount++
		} // not ok result mean that the question is not answered
	}

	score := float32(correctCount) * 100 / float32(questionCount)

	studentQuiz.CorrectCount = &correctCount
	studentQuiz.TotalScore = &score
	studentQuiz.Status = enum.SUBMITTED

	return s.db.Session(&gorm.Session{FullSaveAssociations: true}).Omit("Quiz").Updates(&studentQuiz).Error
}
