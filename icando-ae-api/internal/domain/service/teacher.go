package service

import (
	"icando/internal/domain/repository"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"icando/utils/httperror"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TeacherService interface {
	GetAllTeachers(filter dto.GetTeacherFilter) ([]dao.LearningDesignerDao, *httperror.HttpError)
	FindTeacherByID(id uuid.UUID) (*dao.LearningDesignerDao, *httperror.HttpError)
	PutUserInfo(id uuid.UUID, dto dto.PutUserInfoDto) (*dao.LearningDesignerDao, *httperror.HttpError)
}

type TeacherServiceImpl struct {
	teacherRepository repository.TeacherRepository
	config            *lib.Config
}

func NewTeacherServiceImpl(
	teacherRepository repository.TeacherRepository,
	config *lib.Config,
) *TeacherServiceImpl {
	return &TeacherServiceImpl{
		teacherRepository: teacherRepository,
		config:            config,
	}

}

func (s *TeacherServiceImpl) GetAllTeachers(filter dto.GetTeacherFilter) ([]dao.LearningDesignerDao, *httperror.HttpError) {
	teachers, err := s.teacherRepository.GetAllTeacher(filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeacherNotFound
		}
		return nil, httperror.InternalServerError
	}

	payload := make([]dao.LearningDesignerDao, 0)
	for _, teacher := range teachers {
		payload = append(payload, dao.LearningDesignerDao{
			ID:        teacher.ID,
			FirstName: teacher.FirstName,
			LastName:  teacher.LastName,
			Email:     teacher.Email,
		})
	}

	return payload, nil
}

func (s *TeacherServiceImpl) FindTeacherByID(id uuid.UUID) (
	*dao.LearningDesignerDao, *httperror.HttpError,
) {
	teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeacherNotFound
		}
		return nil, httperror.InternalServerError
	}
	teacherDao := dao.LearningDesignerDao{
		ID:        teacher.ID,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
	}

	return &teacherDao, nil
}

func (s *TeacherServiceImpl) PutUserInfo(
	id uuid.UUID, putUserInfoDto dto.PutUserInfoDto,
) (*dao.LearningDesignerDao, *httperror.HttpError) {
	teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeacherNotFound
		}
		return nil, httperror.InternalServerError
	}
	teacher.FirstName = putUserInfoDto.FirstName
	teacher.LastName = putUserInfoDto.LastName
	teacher.Email = putUserInfoDto.Email
	err = s.teacherRepository.UpdateTeacher(teacher)
	if err != nil {
		return nil, httperror.InternalServerError
	}
	teacherDao := dao.LearningDesignerDao{
		ID:        teacher.ID,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
	}
	return &teacherDao, nil
}
