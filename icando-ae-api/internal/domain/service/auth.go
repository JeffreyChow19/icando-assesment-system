package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"icando/utils/httperror"
	"net/http"
	"time"
)

type AuthService interface {
	Login(loginDto dto.LoginDto, role model.Role) (*dao.AuthDao, error)
	ChangePassword(id uuid.UUID, dto dto.ChangePasswordDto) *httperror.HttpError
}

type AuthServiceImpl struct {
	learningDesignerRepository repository.LearningDesignerRepository
	config                     *lib.Config
}


func NewAuthServiceImpl(learningDesignerRepository repository.LearningDesignerRepository, config *lib.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		learningDesignerRepository: learningDesignerRepository,
		config:                     config,
	}
}

func (s *AuthServiceImpl) Login(loginDto dto.LoginDto, role model.Role) (*dao.AuthDao, error) {
	var authDao *dao.AuthDao 
	if role == model.ROLE_LEARNING_DESIGNER{
		learningDesigner, err := s.learningDesignerRepository.FindUserByEmail(loginDto.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrLearningDesignerNotFound
			}
			return nil, httperror.InternalServerError
		}

		if !s.checkPassword(loginDto.Password, learningDesigner.Password) {
			return nil, httperror.UnauthorizedError
		}

		authDao, err := s.buildAuthDao(learningDesigner)
		if err != nil {
			return nil, httperror.InternalServerError
		}
		return authDao, nil
	}
	return authDao, nil
}

func (s *AuthServiceImpl) ChangePassword(id uuid.UUID, dto dto.ChangePasswordDto) *httperror.HttpError {
	user, err := s.learningDesignerRepository.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLearningDesignerNotFound
		}
		return httperror.InternalServerError
	}

	if !s.checkPassword(dto.OldPassword, user.Password) {
		return httperror.UnauthorizedError
	}

	user.Password = s.hashPassword(dto.NewPassword)
	err = s.learningDesignerRepository.UpdateUserInfo(user)
	if err != nil {
		return httperror.InternalServerError
	}
	return nil
}

// Learning Designer Auth Dao
func (s *AuthServiceImpl) buildAuthDao(learningDesigner *model.LearningDesigner) (*dao.AuthDao, error) {
	token, err := s.generateToken(learningDesigner)
	if err != nil {
		return nil, err
	}
	authDao := dao.AuthDao{
		User: dao.LearningDesignerDao{
			ID:        learningDesigner.ID,
			FirstName: learningDesigner.FirstName,
			LastName:  learningDesigner.LastName,
			Email:     learningDesigner.Email,
		},
		Token: token,
	}

	return &authDao, nil
}

func (s *AuthServiceImpl) generateToken(user *model.LearningDesigner) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(time.Hour * 8).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(s.config.JwtSecret))
	return tokenString, err
}

func (s *AuthServiceImpl) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func (s *AuthServiceImpl) checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var ErrLearningDesignerNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("User not found"),
}
var EmailConflictErr = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("User with that email already exists"),
}
