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

	// Get competencies by IDs
	competencies, err := s.competencyRepository.GetCompetenciesByIDs(questionDto.Competencies)
	if err != nil {
		return nil, ErrCompetencyNotFound
	}

	// Add competencies to the question
	for _, competency := range competencies {
		question.Competencies = append(question.Competencies, competency)
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

	// Get all existing competencies
	existingCompetencies, errGetExistingCompetencies := s.questionCompetencyRepository.GetAll(question.ID)
	if errGetExistingCompetencies != nil {
		return nil, ErrUpdateQuestion
	}
	existingCompetencyMap := make(map[uuid.UUID]bool)
	for _, competency := range existingCompetencies {
		existingCompetencyMap[competency.CompetencyID] = true
	}

	// Get updated competencies
	updatedCompetencies, errGetUpdatedCompetencies := s.competencyRepository.GetCompetenciesByIDs(questionDto.Competencies)
	if errGetUpdatedCompetencies != nil {
		return nil, ErrUpdateQuestion
	}
	updatedCompetencyMap := make(map[uuid.UUID]*model.Competency)
	for i := range updatedCompetencies {
		updatedCompetencyMap[updatedCompetencies[i].ID] = &updatedCompetencies[i]
	}

	// Prepare the list of competencies to be deleted and added
	toBeDeletedCompetencies := []model.QuestionCompetency{}
	toBeAddedCompetencies := []model.Competency{}

	// Find competencies to be deleted and added
	for competencyID := range existingCompetencyMap {
		if _, ok := updatedCompetencyMap[competencyID]; !ok {
			toBeDeletedCompetencies = append(toBeDeletedCompetencies, model.QuestionCompetency{
				QuestionID:   question.ID,
				CompetencyID: competencyID,
			})
		}
	}
	for competencyID, competency := range updatedCompetencyMap {
		if _, ok := existingCompetencyMap[competencyID]; !ok {
			toBeAddedCompetencies = append(toBeAddedCompetencies, *competency)
		}
	}

	// Delete competencies that are no longer part of the question
	err = s.questionCompetencyRepository.Delete(toBeDeletedCompetencies)
	if err != nil {
		return nil, ErrUpdateQuestion
	}

	// Add competencies that are not yet part of the question
	question.Competencies = toBeAddedCompetencies

	// Save the updated question
	err = s.questionRepository.UpdateQuestion(question)
	if err != nil {
		return nil, ErrUpdateQuestion
	}

	question.Competencies = updatedCompetencies

	questionDao, errDao := question.ToDao()
	if errDao != nil {
		return nil, ErrUpdateQuestion
	}

	return questionDao, nil
}
