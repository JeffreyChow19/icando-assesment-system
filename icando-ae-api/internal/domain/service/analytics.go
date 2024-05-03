package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	GetStudentStatistics(studentID uuid.UUID) (*dao.GetStudentStatisticsDao, *httperror.HttpError)
	GetDashboardOverview(id uuid.UUID) (*dao.DashboardOverviewDao, *httperror.HttpError)
}

type AnalyticsServiceImpl struct {
	analyticsRepository repository.AnalyticsRepository
	studentRepository   repository.StudentRepository
}

var ErrGetQuizPerformance = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when fetching quiz performance"),
}

func NewAnalyticsServiceImpl(
	analyticsRepository repository.AnalyticsRepository, studentRepository repository.StudentRepository,
) *AnalyticsServiceImpl {
	return &AnalyticsServiceImpl{
		analyticsRepository: analyticsRepository,
		studentRepository:   studentRepository,
	}
}

func (s *AnalyticsServiceImpl) GetQuizPerformance(filter dto.GetQuizPerformanceFilter) (
	*dao.QuizPerformanceDao, *httperror.HttpError,
) {
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

func (s *AnalyticsServiceImpl) GetLatestSubmissions(filter dto.GetLatestSubmissionsFilter) (
	*[]dao.GetLatestSubmissionsDao, *httperror.HttpError,
) {
	latestSubmissions, err := s.analyticsRepository.GetLatestSubmissions(&filter)

	if err != nil {
		return nil, ErrGetLatestSubmissions
	}

	return latestSubmissions, nil
}

var ErrGetStudentStatistics = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when fetching student statistics"),
}

func (s *AnalyticsServiceImpl) GetStudentStatistics(studentID uuid.UUID) (
	*dao.GetStudentStatisticsDao, *httperror.HttpError,
) {
	studentIDStr := studentID.String()

	// Student Information
	student, errStudent := s.studentRepository.GetOne(dto.GetStudentFilter{ID: &studentIDStr, IncludeClass: true})
	if errStudent != nil {
		if errors.Is(errStudent, gorm.ErrRecordNotFound) {
			return nil, ErrStudentNotFound
		}
		return nil, ErrGetStudentStatistics
	}

	studentDao := student.ToDao()
	classDao := student.Class.ToDao(dto.GetClassFilter{})

	// Performance
	quizPerformance, errQuizPerformance := s.analyticsRepository.GetQuizPerformance(
		&dto.GetQuizPerformanceFilter{
			StudentID: &studentIDStr,
		},
	)
	if errQuizPerformance != nil {
		return nil, ErrGetQuizPerformance
	}

	// Competency
	competencyStats, errCompetencyStats := s.analyticsRepository.GetStudentQuizCompetency(studentID)
	if errCompetencyStats != nil {
		return nil, ErrGetStudentStatistics
	}

	// Quizzes
	quizzes, errQuizzes := s.analyticsRepository.GetStudentQuizzes(studentID)
	if errQuizzes != nil {
		return nil, ErrGetStudentStatistics
	}

	return &dao.GetStudentStatisticsDao{
		StudentInfo: dao.StudentInfo{
			Student: studentDao,
			Class:   classDao,
		},
		Performance: *quizPerformance,
		Competency:  *competencyStats,
		Quizzes:     *quizzes,
	}, nil
}

func (s *AnalyticsServiceImpl) GetDashboardOverview(id uuid.UUID) (*dao.DashboardOverviewDao, *httperror.HttpError) {
	dashboardDao, err := s.analyticsRepository.GetTeacherDashboardOverview(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeacherNotFound
		}
		return nil, httperror.InternalServerError
	}
	return dashboardDao, nil
}
