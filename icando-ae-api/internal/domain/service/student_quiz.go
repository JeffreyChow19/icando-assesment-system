package service

import (
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/internal/worker/client"
	"icando/lib"
	"icando/utils/httperror"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type StudentQuizService interface {
	StartQuiz(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError)
	SubmitQuiz(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError)
	UpdateStudentAnswer(studentQuiz *model.StudentQuiz, questionID uuid.UUID, studentAnswerDto dto.UpdateStudentAnswerDto) *httperror.HttpError
	CalculateScore(id uuid.UUID) error
	GetQuizAvailability(studentQuiz *model.StudentQuiz) (*dao.QuizDao, *httperror.HttpError)
	GetQuizDetail(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError)
}

type StudentQuizServiceImpl struct {
	studentQuizRepository repository.StudentQuizRepository
	questionRepository    repository.QuestionRepository
	quizRepository        repository.QuizRepository
	db                    *gorm.DB
	workerClient          *client.WorkerClient
}

func NewStudentQuizServiceImpl(
	studentQuizRepository repository.StudentQuizRepository,
	questionRepository repository.QuestionRepository,
	quizRepository repository.QuizRepository,
	workerClient *client.WorkerClient,
	db *lib.Database) *StudentQuizServiceImpl {
	return &StudentQuizServiceImpl{
		studentQuizRepository: studentQuizRepository,
		questionRepository:    questionRepository,
		quizRepository:        quizRepository,
		db:                    db.DB,
		workerClient:          workerClient,
	}
}

var ErrStartQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Error on starting quiz"),
}
var ErrQuizStarted = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("Quiz has already been started"),
}

func (s *StudentQuizServiceImpl) StartQuiz(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError) {
	if studentQuiz.Status == enum.SUBMITTED {
		return nil, ErrStudentQuizSubmitted
	}

	if studentQuiz.Status != enum.NOT_STARTED || studentQuiz.StartedAt != nil {
		return nil, ErrQuizStarted
	}

	currentTime := time.Now()
	studentQuiz.StartedAt = &currentTime
	studentQuiz.Status = enum.STARTED

	err := s.studentQuizRepository.UpdateStudentQuiz(*studentQuiz)
	if err != nil {
		return nil, ErrStartQuiz
	}

	resp, errResp := studentQuiz.ToDao(false)
	if errResp != nil {
		return nil, ErrStartQuiz
	}

	return resp, nil
}

var ErrSubmitQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Error to submit quiz"),
}

func (s *StudentQuizServiceImpl) SubmitQuiz(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError) {
	if studentQuiz.Status == enum.NOT_STARTED {
		return nil, ErrStudentQuizNotStarted
	}

	if studentQuiz.Status == enum.SUBMITTED {
		return nil, ErrStudentQuizSubmitted
	}

	currentTime := time.Now()
	studentQuiz.CompletedAt = &currentTime
	studentQuiz.Status = enum.SUBMITTED

	err := s.CalculateScore(studentQuiz.ID)

	if err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	resp, errResp := studentQuiz.ToDao(false)
	if errResp != nil {
		return nil, ErrSubmitQuiz
	}

	return resp, nil
}

var ErrInvalidQuestion = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invalid question"),
}
var ErrStudentQuizNotStarted = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Student quiz has not been started"),
}
var ErrStudentQuizSubmitted = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Student quiz has already been submitted"),
}
var ErrQuizHasNotStarted = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Quiz has not started yet"),
}
var ErrQuizHasEnded = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Quiz has ended"),
}
var ErrQuizDurationHasEnded = &httperror.HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Quiz duration has ended"),
}
var ErrUpdateStudentAnswer = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating student answer"),
}
var ErrGetQuizDetail = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when getting quiz detail"),
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
	if studentQuiz.StartedAt != nil && currentTime.After(studentQuiz.StartedAt.Add(time.Minute*time.Duration(*quiz.Duration))) {
		return ErrQuizDurationHasEnded
	}
	if quiz.EndAt != nil && currentTime.After(*quiz.EndAt) {
		return ErrQuizHasEnded
	}
	if quiz.StartAt != nil && currentTime.Before(*quiz.StartAt) {
		return ErrQuizHasNotStarted
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

	if studentQuiz.Status != enum.STARTED {
		return nil // skip not started and submitted
	}

	answerMap := make(map[string]model.StudentAnswer)
	competencyMap := make(map[string]dto.StudentQuizCompetencyCorrectTotalDto)

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

				comp, ok := competencyMap[competency.ID.String()]

				if ok {
					comp.TotalCount++

					if isCorrect {
						comp.CorrectCount++
					}

					competencyMap[competency.ID.String()] = comp
				} else {
					if isCorrect {
						competencyMap[competency.ID.String()] = dto.StudentQuizCompetencyCorrectTotalDto{
							TotalCount:   1,
							CorrectCount: 1,
						}
					} else {
						competencyMap[competency.ID.String()] = dto.StudentQuizCompetencyCorrectTotalDto{
							TotalCount:   1,
							CorrectCount: 0,
						}
					}
				}
			}

			if err := answer.SetCompetencies(answerCompetencies); err != nil {
				return err
			}

			answerMap[question.ID.String()] = answer
			correctCount++
		} // not ok result mean that the question is not answered
	}

	updatedAnswers := make([]model.StudentAnswer, 0)

	for _, answer := range answerMap {
		updatedAnswers = append(updatedAnswers, answer)
	}

	score := float32(correctCount) * 100 / float32(questionCount)

	studentQuiz.StudentAnswers = updatedAnswers
	studentQuiz.CorrectCount = &correctCount
	studentQuiz.TotalScore = &score
	studentQuiz.Status = enum.SUBMITTED

	now := time.Now()
	studentQuiz.CompletedAt = &now

	studentQuizCompetencies := make([]model.StudentQuizCompetency, 0)

	for id, sqc := range competencyMap {
		studentQuizCompetencies = append(studentQuizCompetencies, model.StudentQuizCompetency{
			StudentID:     studentQuiz.StudentID,
			StudentQuizID: studentQuiz.ID,
			CompetencyID:  uuid.MustParse(id),
			TotalCount:    sqc.TotalCount,
			CorrectCount:  sqc.CorrectCount,
		})
	}

	tx := s.db.Begin()

	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Omit("Quiz").Updates(&studentQuiz).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(studentQuizCompetencies) != 0 {
		if err := tx.Create(&studentQuizCompetencies).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *StudentQuizServiceImpl) GetQuizAvailability(studentQuiz *model.StudentQuiz) (*dao.QuizDao, *httperror.HttpError) {
	quiz, err := s.quizRepository.GetQuiz(dto.GetQuizFilter{ID: studentQuiz.QuizID})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrGetQuiz
	}

	newQuizCount, err := s.quizRepository.CheckNewQuizVersion(studentQuiz.QuizID, studentQuiz.StudentID)

	if err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Failed to check new quiz version availability"),
		}
	}

	hasNewerVersion := *newQuizCount > 0
	quiz.HasNewerVersion = &hasNewerVersion

	quizDao := quiz.ToDao(false)

	return &quizDao, nil
}

func (s *StudentQuizServiceImpl) GetQuizDetail(studentQuiz *model.StudentQuiz) (*dao.StudentQuizDao, *httperror.HttpError) {
	if studentQuiz.Status == enum.NOT_STARTED {
		return nil, ErrStudentQuizNotStarted
	}

	if studentQuiz.Status == enum.SUBMITTED {
		return nil, ErrStudentQuizSubmitted
	}

	studentQuiz, err := s.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{ID: studentQuiz.ID,
		WithQuizQuestions: true,
		WithAnswers:       true,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrGetQuiz
	}

	currentTime := time.Now()

	if studentQuiz.StartedAt != nil && currentTime.After(studentQuiz.StartedAt.Add(time.Minute*time.Duration(*studentQuiz.Quiz.Duration))) {
		return nil, ErrQuizDurationHasEnded
	}

	if studentQuiz.Quiz.StartAt != nil && currentTime.Before(*studentQuiz.Quiz.StartAt) {
		return nil, ErrQuizHasNotStarted
	}

	if studentQuiz.Quiz.EndAt != nil && currentTime.After(*studentQuiz.Quiz.EndAt) {
		return nil, ErrQuizHasEnded
	}

	quizDao, err := studentQuiz.ToDao(false)
	if err != nil {
		return nil, ErrGetQuizDetail
	}

	return quizDao, nil
}
