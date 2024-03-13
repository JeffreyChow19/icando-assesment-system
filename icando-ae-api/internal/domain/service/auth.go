package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/lib"
	"icando/utils/httperror"
	"net/http"
	"time"
)

type AuthService interface {
	Login(loginDto dto.LoginDto, role enum.Role) (*dao.AuthDao, *httperror.HttpError)
	ChangePassword(id uuid.UUID, role enum.Role, dto dto.ChangePasswordDto) *httperror.HttpError
	ProfileStudent(id uuid.UUID) (*dao.StudentDao, *httperror.HttpError)
	ProfileTeacher(id uuid.UUID) (*dao.TeacherDao, *httperror.HttpError)
}

type AuthServiceImpl struct {
	studentRepository repository.StudentRepository
	teacherRepository repository.TeacherRepository
	config            *lib.Config
}

func NewAuthServiceImpl(
	studentRepository repository.StudentRepository,
	teacherRepository repository.TeacherRepository,
	config *lib.Config,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		studentRepository: studentRepository,
		teacherRepository: teacherRepository,
		config:            config,
	}
}

func (s *AuthServiceImpl) Login(loginDto dto.LoginDto, role enum.Role) (*dao.AuthDao, *httperror.HttpError) {
	if role == enum.ROLE_STUDENT {
		// todo student token should have different expire date (1 week or manually configurable)
		student, err := s.studentRepository.GetOne(dto.GetStudentFilter{Email: &loginDto.Email})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrStudentNotFound
			}
			return nil, httperror.InternalServerError
		}

		claim := dao.TokenClaim{
			ID: student.ID,
		}

		authDao, err := s.buildAuthDao(claim)

		if err != nil {
			return nil, httperror.InternalServerError
		}

		return authDao, nil
	} else if role == enum.ROLE_TEACHER || role == enum.ROLE_LEARNING_DESIGNER {
		teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{Email: &loginDto.Email})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrTeacherNotFound
			}
			return nil, httperror.InternalServerError
		}

		if teacher.Role != enum.ROLE_LEARNING_DESIGNER && role == enum.ROLE_LEARNING_DESIGNER {
			return nil, httperror.UnauthorizedError
		}

		if !s.checkPassword(loginDto.Password, teacher.Password) {
			return nil, httperror.UnauthorizedError
		}

		claim := dao.TokenClaim{
			ID:  teacher.ID,
			Exp: time.Now().Add(time.Hour * 8).Unix(),
		}

		authDao, err := s.buildAuthDao(claim)

		if err != nil {
			return nil, httperror.InternalServerError
		}

		return authDao, nil
	}

	return nil, httperror.InternalServerError
}

func (s *AuthServiceImpl) ChangePassword(
	id uuid.UUID, role enum.Role, changePasswordDto dto.ChangePasswordDto,
) *httperror.HttpError {
	if role == enum.ROLE_TEACHER || role == enum.ROLE_LEARNING_DESIGNER {
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
	// to do: handle for student

	return nil
}

func (s *AuthServiceImpl) ProfileStudent(id uuid.UUID) (*dao.StudentDao, *httperror.HttpError) {
	idString := id.String()
	student, err := s.studentRepository.GetOne(dto.GetStudentFilter{ID: &idString})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentNotFound
		}
		return nil, ErrCreateStudent
	}

	studentDao := student.ToDao()

	return &studentDao, nil
}

func (s *AuthServiceImpl) ProfileTeacher(id uuid.UUID) (*dao.TeacherDao, *httperror.HttpError) {
	teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &id})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTeacherNotFound
		}
		return nil, httperror.InternalServerError
	}

	teacherDao := teacher.ToDao()

	return &teacherDao, nil
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
			"id":  user.ID.String(),
			"exp": user.Exp,
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

var ErrTeacherNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Teacher not found"),
}

var EmailConflictErr = &httperror.HttpError{
	StatusCode: http.StatusConflict,
	Err:        errors.New("User with that email already exists"),
}

var InvalidCredentialsError = &httperror.HttpError{
	StatusCode: http.StatusUnauthorized,
	Err:        errors.New("Invalid email or password credentials"),
}
