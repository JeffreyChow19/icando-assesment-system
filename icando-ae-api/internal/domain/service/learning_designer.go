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
	FindUserById(id uuid.UUID) (*dao.LearningDesignerDao, *httperror.HttpError)
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
func (s *LearningDesignerServiceImpl) FindUserById(id uuid.UUID) (*dao.LearningDesignerDao, *httperror.HttpError) {
	user, err := s.learningDesignerRepository.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundErr
		}
		return nil, httperror.InternalServerError
	}
	userDao := dao.LearningDesignerDao{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      "Learning Designer",
	}

	return &userDao, nil
}

func (s *LearningDesignerServiceImpl) PutUserInfo(id uuid.UUID, dto dto.PutUserInfoDto) (*dao.LearningDesignerDao, *httperror.HttpError) {
	user, err := s.learningDesignerRepository.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundErr
		}
		return nil, httperror.InternalServerError
	}
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Email = dto.Email
	err = s.learningDesignerRepository.UpdateUserInfo(user)
	if err != nil {
		return nil, httperror.InternalServerError
	}
	userDao := dao.LearningDesignerDao{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      "Learning Designer",
	}
	return &userDao, nil
}
