package service

import (
	"github.com/google/uuid"
	"icando/internal/domain/repository"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
)

type ClassService interface {
	GetAllClass(filter dto.GetAllClassFilter) ([]dao.ClassDao, error)
	GetClass(classID uuid.UUID, fitler dto.GetClassFitler) (*dao.ClassDao, error)
	CreateClass(classDto dto.ClassDto) (*dao.ClassDao, error)
	UpdateClass(id uuid.UUID, classDto dto.ClassDto) (*dao.ClassDao, error)
	DeleteClass(id uuid.UUID) error
}

type ClassServiceImpl struct {
	classRepository repository.ClassRepository
}

func NewClassServiceImpl(classRepository repository.ClassRepository) *ClassServiceImpl {
	return &ClassServiceImpl{
		classRepository: classRepository,
	}
}

func (s *ClassServiceImpl) GetAllClass(filter dto.GetAllClassFilter) ([]dao.ClassDao, error) {
	class, err := s.classRepository.GetAllClass(filter)
	if err != nil {
		return nil, err
	}

	payload := make([]dao.ClassDao, 0)

	for _, cls := range class {
		payload = append(payload, cls.ToDao(dto.GetClassFitler{}))
	}

	return payload, nil
}

func (s *ClassServiceImpl) GetClass(classID uuid.UUID, filter dto.GetClassFitler) (*dao.ClassDao, error) {
	class, err := s.classRepository.GetClass(classID, filter)

	if err != nil {
		return nil, err
	}

	classDao := class.ToDao(filter)

	return &classDao, nil
}

func (s *ClassServiceImpl) CreateClass(classDto dto.ClassDto) (*dao.ClassDao, error) {
	class, err := s.classRepository.CreateClass(classDto)

	if err != nil {
		return nil, err
	}

	classDao := class.ToDao(dto.GetClassFitler{})

	return &classDao, nil
}

func (s *ClassServiceImpl) UpdateClass(id uuid.UUID, classDto dto.ClassDto) (*dao.ClassDao, error) {
	class, err := s.classRepository.UpdateClass(id, classDto)

	if err != nil {
		return nil, err
	}

	classDao := class.ToDao(dto.GetClassFitler{})

	return &classDao, nil
}

func (s *ClassServiceImpl) DeleteClass(id uuid.UUID) error {
	err := s.classRepository.DeleteClass(id)
	return err
}
