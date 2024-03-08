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
	studentRepository          repository.StudentRepository
	teacherRepository          repository.TeacherRepository
	config                     *lib.Config
}

func NewAuthServiceImpl(learningDesignerRepository repository.LearningDesignerRepository,
	studentRepository repository.StudentRepository,
	teacherRepository repository.TeacherRepository,
	config *lib.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		learningDesignerRepository: learningDesignerRepository,
		studentRepository:          studentRepository,
		teacherRepository:          teacherRepository,
		config:                     config,
	}
}

func (s *AuthServiceImpl) Login(loginDto dto.LoginDto, role model.Role) (*dao.AuthDao, error) {
	if role == model.ROLE_LEARNING_DESIGNER {
		learningDesigner, err := s.learningDesignerRepository.FindLearningDesigner(dto.GetLearningDesignerFilter{Email: &loginDto.Email})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrLearningDesignerNotFound
			}
			return nil, httperror.InternalServerError
		}

		if !s.checkPassword(loginDto.Password, learningDesigner.Password) {
			return nil, httperror.UnauthorizedError
		}

		claim := dao.TokenClaim{
			ID:   learningDesigner.ID,
			Role: model.ROLE_LEARNING_DESIGNER,
		}

		authDao, err := s.buildAuthDao(claim)

		if err != nil {
			return nil, httperror.InternalServerError
		}

		return authDao, nil
	} else if role == model.ROLE_STUDENT {
		// todo student token should have different expire date (1 week or manually configurable)
		student, err := s.studentRepository.GetOne(dto.GetStudentFilter{Email: &loginDto.Email})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrStudentNotFound
			}
			return nil, httperror.InternalServerError
		}

		claim := dao.TokenClaim{
			ID:   student.ID,
			Role: model.ROLE_STUDENT,
		}

		authDao, err := s.buildAuthDao(claim)

		if err != nil {
			return nil, httperror.InternalServerError
		}

		return authDao, nil
	} else if role == model.ROLE_TEACHER {
		teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{Email: &loginDto.Email})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrTeacherNotFound
			}
			return nil, httperror.InternalServerError
		}

		if !s.checkPassword(loginDto.Password, teacher.Password) {
			return nil, httperror.UnauthorizedError
		}

		claim := dao.TokenClaim{
			ID:   teacher.ID,
			Role: model.ROLE_TEACHER,
		}

		authDao, err := s.buildAuthDao(claim)

		if err != nil {
			return nil, httperror.InternalServerError
		}

		return authDao, nil
	}

	return nil, errors.New("Unhandled data")
}

func (s *AuthServiceImpl) ChangePassword(id uuid.UUID, role model.Role, changePasswordDto dto.ChangePasswordDto) *httperror.HttpError {
	if role == model.ROLE_LEARNING_DESIGNER {
		user, err := s.learningDesignerRepository.FindLearningDesigner(dto.GetLearningDesignerFilter{ID: &id})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrLearningDesignerNotFound
			}
			return httperror.InternalServerError
		}

		if !s.checkPassword(changePasswordDto.OldPassword, user.Password) {
			return httperror.UnauthorizedError
		}

		user.Password = s.hashPassword(changePasswordDto.NewPassword)
		err = s.learningDesignerRepository.UpdateUserInfo(user)
		if err != nil {
			return httperror.InternalServerError
		}
	} else if role == model.ROLE_TEACHER {
		user, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &id})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrTeacherNotFound
			}
			return httperror.InternalServerError
		}

		if !s.checkPassword(changePasswordDto.OldPassword, user.Password) {
			return httperror.UnauthorizedError
		}

		user.Password = s.hashPassword(changePasswordDto.NewPassword)
		err = s.teacherRepository.UpdateTeacher(user)
		if err != nil {
			return httperror.InternalServerError
		}
	}

	return nil
}

func (s *AuthServiceImpl) buildAuthDao(claim dao.TokenClaim) (*dao.AuthDao, error) {

	token, err := s.generateToken(claim)
	if err != nil {
		return nil, err
	}

	authDao := dao.AuthDao{
		User:  claim,
		Token: token,
	}

	return &authDao, nil
}

func (s *AuthServiceImpl) generateToken(user dao.TokenClaim) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   user.ID.String(),
			"role": string(user.Role),
			"exp":  time.Now().Add(time.Hour * 8).Unix(),
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
	Err:        errors.New("Learning Designer not found"),
}

var ErrTeacherNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Teacher not found"),
}

var EmailConflictErr = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("User with that email already exists"),
}
