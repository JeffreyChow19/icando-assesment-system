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
	RegisterUser(registerDto *dto.RegisterDto) (*dao.AuthDao, *httperror.HttpError)
	Login(loginDro dto.LoginDto) (*dao.AuthDao, error)
	ChangePassword(id uuid.UUID, dto dto.ChangePasswordDto) *httperror.HttpError
}

type AuthServiceImpl struct {
	learningDesignerRepository repository.LearningDesignerRepository
	config                     *lib.Config
}

func (s *AuthServiceImpl) RegisterUser(registerDto *dto.RegisterDto) (*dao.AuthDao, *httperror.HttpError) {
	if _, err := s.learningDesignerRepository.FindUserByEmail(registerDto.Email); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httperror.InternalServerError
		}
	} else {
		return nil, EmailConflictErr
	}

	hashedPassword := s.hashPassword(registerDto.Password)
	user := model.LearningDesigner{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Password:  hashedPassword,
		Email:     registerDto.Email,
	}
	user.ID = uuid.New()

	err := s.learningDesignerRepository.AddUser(&user)
	if err != nil {
		return nil, httperror.InternalServerError
	}
	authDao, err := s.buildAuthDao(&user)
	if err != nil {
		return nil, httperror.InternalServerError
	}

	return authDao, nil
}

func NewAuthServiceImpl(learningDesignerRepository repository.LearningDesignerRepository, config *lib.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		learningDesignerRepository: learningDesignerRepository,
		config:                     config,
	}
}

func (s *AuthServiceImpl) Login(loginDto dto.LoginDto) (*dao.AuthDao, error) {
	user, err := s.learningDesignerRepository.FindUserByEmail(loginDto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundErr
		}
		return nil, httperror.InternalServerError
	}

	if !s.checkPassword(loginDto.Password, user.Password) {
		return nil, httperror.UnauthorizedError
	}

	authDao, err := s.buildAuthDao(user)
	if err != nil {
		return nil, httperror.InternalServerError
	}

	return authDao, nil
}

func (s *AuthServiceImpl) ChangePassword(id uuid.UUID, dto dto.ChangePasswordDto) *httperror.HttpError {
	user, err := s.learningDesignerRepository.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserNotFoundErr
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

func (s *AuthServiceImpl) buildAuthDao(user *model.LearningDesigner) (*dao.AuthDao, error) {
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}
	authDao := dao.AuthDao{
		User: dao.LearningDesignerDao{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			// Role:        "Learning Designer",
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

var UserNotFoundErr = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("User not found"),
}
var EmailConflictErr = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("User with that email already exists"),
}
