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

type CompetencyService interface {
	GetAllCompetencies(filter dto.GetAllCompetenciesFilter) ([]dao.CompetencyDao, *dao.MetaDao, *httperror.HttpError)
	CreateCompetency(competency dto.CreateCompetencyDto) (*dao.CompetencyDao, *httperror.HttpError)
	UpdateCompetency(competency dto.UpdateCompetencyDto) (*dao.CompetencyDao, *httperror.HttpError)
	DeleteCompetency(competencyId uuid.UUID) (*dao.CompetencyDao, *httperror.HttpError)
}

type CompetencyServiceImpl struct {
	competencyRepository repository.CompetencyRepository
}

func NewCompetencyServiceImpl(competencyRepository repository.CompetencyRepository) *CompetencyServiceImpl {
	return &CompetencyServiceImpl{
		competencyRepository: competencyRepository,
	}
}

var ErrGetAllCompetencies = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when getting all competencies"),
}

func (s *CompetencyServiceImpl) GetAllCompetencies(filter dto.GetAllCompetenciesFilter) ([]dao.CompetencyDao, *dao.MetaDao, *httperror.HttpError) {
	competencies, meta, err := s.competencyRepository.GetAllCompetencies(filter)
	if err != nil {
		return nil, nil, ErrGetAllCompetencies
	}

	var competenciesDao []dao.CompetencyDao
	for _, competency := range competencies {
		competenciesDao = append(competenciesDao, competency.ToDao())
	}

	return competenciesDao, meta, nil
}

var ErrDuplicateCompetencyNumbering = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("Numbering has existed and violate unique constraint"),
}

var ErrCreateCompetency = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating competency"),
}

func (s *CompetencyServiceImpl) CreateCompetency(competencyDto dto.CreateCompetencyDto) (*dao.CompetencyDao, *httperror.HttpError) {
	existing, err := s.competencyRepository.GetCompetencyByNumbering(competencyDto.Numbering)
	if err != nil {
		return nil, ErrCreateCompetency
	}
	if existing != nil {
		return nil, ErrDuplicateCompetencyNumbering
	}

	competency := model.Competency{
		Numbering:   competencyDto.Numbering,
		Name:        competencyDto.Name,
		Description: competencyDto.Description,
	}

	err = s.competencyRepository.CreateCompetency(competency)
	if err != nil {
		return nil, ErrCreateCompetency
	}

	competencyDao := competency.ToDao()

	return &competencyDao, nil
}

var ErrCompetencyNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Competency Not Found"),
}

var ErrUpdateCompetency = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating competency"),
}

func (s *CompetencyServiceImpl) UpdateCompetency(competencyDto dto.UpdateCompetencyDto) (*dao.CompetencyDao, *httperror.HttpError) {
	if competencyDto.Numbering != nil {
		existing, err := s.competencyRepository.GetCompetencyByNumbering(*competencyDto.Numbering)
		if err != nil {
			return nil, ErrUpdateCompetency
		}
		if existing != nil {
			return nil, ErrDuplicateCompetencyNumbering
		}
	}

	competency, err := s.competencyRepository.GetCompetencyById(competencyDto.ID)
	if err != nil {
		if errors.Is(err, repository.ErrCompetencyNotFound) {
			return nil, ErrCompetencyNotFound
		}
		return nil, ErrUpdateCompetency
	}

	if competencyDto.Numbering != nil {
		competency.Numbering = *competencyDto.Numbering
	}
	if competencyDto.Name != nil {
		competency.Name = *competencyDto.Name
	}
	if competencyDto.Description != nil {
		competency.Description = *competencyDto.Description
	}

	err = s.competencyRepository.UpdateCompetency(*competency)
	if err != nil {
		return nil, ErrUpdateCompetency
	}

	competencyDao := competency.ToDao()

	return &competencyDao, nil
}

var ErrDeleteCompetency = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when deleting competency"),
}

func (s *CompetencyServiceImpl) DeleteCompetency(competencyId uuid.UUID) (*dao.CompetencyDao, *httperror.HttpError) {
	competency, err := s.competencyRepository.GetCompetencyById(competencyId)
	if err != nil {
		if errors.Is(err, repository.ErrCompetencyNotFound) {
			return nil, ErrCompetencyNotFound
		}
		return nil, ErrDeleteCompetency
	}

	err = s.competencyRepository.DeleteCompetency(*competency)
	if err != nil {
		return nil, ErrDeleteCompetency
	}

	competencyDao := competency.ToDao()

	return &competencyDao, nil
}
