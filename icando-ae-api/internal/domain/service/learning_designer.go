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

type LearningDesignerService interface {
	FindLearningDesignerById(id uuid.UUID) (*dao.LearningDesignerDao, *httperror.HttpError)
	PutUserInfo(id uuid.UUID, dto dto.PutUserInfoDto) (*dao.LearningDesignerDao, *httperror.HttpError)
}

type LearningDesignerServiceImpl struct {
	learningDesignerRepository repository.LearningDesignerRepository
	config                     *lib.Config
}

func NewLearningDesignerServiceImpl(learningDesignerRepository repository.LearningDesignerRepository, config *lib.Config) *LearningDesignerServiceImpl {
	return &LearningDesignerServiceImpl{
		learningDesignerRepository: learningDesignerRepository,
		config:                     config,
	}

}

func (s *LearningDesignerServiceImpl) FindLearningDesignerById(id uuid.UUID) (*dao.LearningDesignerDao, *httperror.HttpError) {
	user, err := s.learningDesignerRepository.FindLearningDesigner(dto.GetLearningDesignerFilter{ID: &id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLearningDesignerNotFound
		}
		return nil, httperror.InternalServerError
	}
	userDao := dao.LearningDesignerDao{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return &userDao, nil
}

func (s *LearningDesignerServiceImpl) PutUserInfo(id uuid.UUID, putUserInfoDto dto.PutUserInfoDto) (*dao.LearningDesignerDao, *httperror.HttpError) {
	user, err := s.learningDesignerRepository.FindLearningDesigner(dto.GetLearningDesignerFilter{ID: &id})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLearningDesignerNotFound
		}
		return nil, httperror.InternalServerError
	}
	user.FirstName = putUserInfoDto.FirstName
	user.LastName = putUserInfoDto.LastName
	user.Email = putUserInfoDto.Email
	err = s.learningDesignerRepository.UpdateUserInfo(user)
	if err != nil {
		return nil, httperror.InternalServerError
	}
	userDao := dao.LearningDesignerDao{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return &userDao, nil
}
