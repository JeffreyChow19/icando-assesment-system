package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type QuestionService interface {
	CreateQuestion(quizID uuid.UUID, questionDto dto.QuestionDto) (*dao.QuestionDao, *httperror.HttpError)
	UpdateQuestion(filter dto.GetQuestionFilter, questionDto dto.QuestionDto) (*dao.QuestionDao, *httperror.HttpError)
}

type QuestionServiceImpl struct {
	questionRepository           repository.QuestionRepository
	competencyRepository         repository.CompetencyRepository
	questionCompetencyRepository repository.QuestionCompetencyRepository
}

func NewQuestionServiceImpl(questionRepository repository.QuestionRepository, competencyRepository repository.CompetencyRepository, questionCompetencyRepository repository.QuestionCompetencyRepository) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		questionRepository:           questionRepository,
		competencyRepository:         competencyRepository,
		questionCompetencyRepository: questionCompetencyRepository,
	}
}

var ErrCreateQuestion = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating question"),
}

func (s *QuestionServiceImpl) CreateQuestion(quizID uuid.UUID, questionDto dto.QuestionDto) (*dao.QuestionDao, *httperror.HttpError) {
	question := model.Question{
		Text:     questionDto.Text,
		AnswerID: questionDto.AnswerID,
		QuizID:   quizID,
	}
	err := question.SetQuestionChoices(questionDto.Choices)
	if err != nil {
		return nil, ErrCreateQuestion
	}

	// Add competencies to the question
	for _, competencyID := range questionDto.Competencies {
		competency, err := s.competencyRepository.GetOneCompetency(dto.GetOneCompetencyFilter{
			Id: competencyID,
		})
		if err != nil {
			return nil, ErrCompetencyNotFound
		}
		question.Competencies = append(question.Competencies, *competency)
	}

	question, err = s.questionRepository.CreateQuestion(question)
	if err != nil {
		return nil, ErrCreateQuestion
	}

	questionDao, errDao := question.ToDao()
	if errDao != nil {
		return nil, ErrCreateQuestion
	}

	return questionDao, nil
}

var ErrQuestionNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Question Not Found"),
}

var ErrUpdateQuestion = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating question"),
}

func (s *QuestionServiceImpl) UpdateQuestion(filter dto.GetQuestionFilter, questionDto dto.QuestionDto) (*dao.QuestionDao, *httperror.HttpError) {
	// Get the existing question
	question, err := s.questionRepository.GetQuestion(filter)
	if err != nil {
		return nil, ErrQuestionNotFound
	}

	// Update the question fields
	question.Text = questionDto.Text
	question.AnswerID = questionDto.AnswerID
	err = question.SetQuestionChoices(questionDto.Choices)
	if err != nil {
		return nil, ErrUpdateQuestion
	}

	// Delete all existing competencies in the question_competency relation
	err = s.questionCompetencyRepository.Delete(question.ID)
	if err != nil {
		return nil, ErrUpdateQuestion
	}

	// Update competencies
	question.Competencies = []model.Competency{}
	for _, competencyID := range questionDto.Competencies {
		competency, err := s.competencyRepository.GetOneCompetency(dto.GetOneCompetencyFilter{
			Id: competencyID,
		})
		if err != nil {
			return nil, ErrCompetencyNotFound
		}
		question.Competencies = append(question.Competencies, *competency)
	}

	// Save the updated question
	err = s.questionRepository.UpdateQuestion(*question)
	if err != nil {
		return nil, ErrUpdateQuestion
	}

	questionDao, errDao := question.ToDao()
	if errDao != nil {
		return nil, ErrUpdateQuestion
	}

	return questionDao, nil
}
