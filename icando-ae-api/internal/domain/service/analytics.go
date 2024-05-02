package service

import (
	"icando/internal/domain/repository"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"

	"github.com/pkg/errors"
)

type AnalyticsService interface {
	GetQuizPerformance(filter dto.GetQuizPerformanceFilter) (*dao.QuizPerformanceDao, *httperror.HttpError)
	GetLatestSubmissions(filter dto.GetLatestSubmissionsFilter) (*[]dao.GetLatestSubmissionsDao, *httperror.HttpError)
}

type AnalyticsServiceImpl struct {
	analyticsRepository repository.AnalyticsRepository
}

var ErrGetQuizPerformance = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when fetching quiz performance"),
}

func NewAnalyticsServiceImpl(analyticsRepository repository.AnalyticsRepository) *AnalyticsServiceImpl {
	return &AnalyticsServiceImpl{
		analyticsRepository: analyticsRepository,
	}
}

func (s *AnalyticsServiceImpl) GetQuizPerformance(filter dto.GetQuizPerformanceFilter) (*dao.QuizPerformanceDao, *httperror.HttpError) {
	quizPerformance, err := s.analyticsRepository.GetQuizPerformance(&filter)

	if err != nil {
		return nil, ErrGetQuizPerformance
	}

	return quizPerformance, nil
}

var ErrGetLatestSubmissions = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when fetching latest submissions"),
}

func (s *AnalyticsServiceImpl) GetLatestSubmissions(filter dto.GetLatestSubmissionsFilter) (*[]dao.GetLatestSubmissionsDao, *httperror.HttpError) {
	latestSubmissions, err := s.analyticsRepository.GetLatestSubmissions(&filter)

	if err != nil {
		return nil, ErrGetLatestSubmissions
	}

	return latestSubmissions, nil
}
