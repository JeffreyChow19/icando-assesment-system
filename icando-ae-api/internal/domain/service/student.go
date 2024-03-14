package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type StudentService interface {
	AddStudent(institutionId uuid.UUID, studentDto dto.CreateStudentDto) (*dao.StudentDao, *httperror.HttpError)
	GetStudent(institutionId uuid.UUID, id uuid.UUID) (*dao.StudentDao, *httperror.HttpError)
	UpdateStudent(institutionId uuid.UUID, id uuid.UUID, dto dto.UpdateStudentDto) (
		*dao.StudentDao, *httperror.HttpError,
	)
	DeleteStudent(institutionId uuid.UUID, id uuid.UUID) *httperror.HttpError
	GetAllStudents(filter dto.GetAllStudentsFilter) (
		[]dao.StudentDao,
		*dao.MetaDao, *httperror.HttpError,
	)
}

type StudentServiceImpl struct {
	studentRepository repository.StudentRepository
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		studentRepository: studentRepository,
	}
}

func (s *StudentServiceImpl) GetAllStudents(filter dto.GetAllStudentsFilter) (
	[]dao.StudentDao,
	*dao.MetaDao, *httperror.HttpError,
) {
	students, meta, err := s.studentRepository.GetAllStudent(filter)
	if err != nil {
		log.Print(err)
		return nil, nil, httperror.InternalServerError
	}

	studentsDao := []dao.StudentDao{}
	for _, student := range students {
		studentsDao = append(studentsDao, student.ToDao())
	}

	return studentsDao, meta, nil
}

func (s *StudentServiceImpl) AddStudent(institutionId uuid.UUID, studentDto dto.CreateStudentDto) (
	*dao.StudentDao, *httperror.HttpError,
) {
	httperr := s.checkStudentUniqueIdentifier(studentDto.Email, studentDto.Nisn)
	if httperr != nil {
		return nil, httperr
	}

	student := model.Student{
		FirstName:     studentDto.FirstName,
		LastName:      studentDto.LastName,
		Email:         studentDto.Email,
		Nisn:          studentDto.Nisn,
		ClassID:       studentDto.ClassID,
		InstitutionID: institutionId,
	}
	err := s.studentRepository.Create(&student)
	if err != nil {
		return nil, ErrCreateStudent
	}
	dao := student.ToDao()
	return &dao, nil
}

func (s *StudentServiceImpl) checkStudentUniqueIdentifier(email string, nisn string) *httperror.HttpError {
	_, err := s.studentRepository.GetOne(dto.GetStudentFilter{Nisn: &nisn})
	if err == nil {
		return ErrDuplicateNisn
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return httperror.InternalServerError
	}

	_, err = s.studentRepository.GetOne(dto.GetStudentFilter{Email: &email})
	if err == nil {
		return ErrDuplicateEmail
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return httperror.InternalServerError
	}

	return nil
}

func (s *StudentServiceImpl) GetStudent(institutionId uuid.UUID, id uuid.UUID) (
	*dao.StudentDao, *httperror.HttpError,
) {
	student, err := s.getStudentById(institutionId, id)
	if err != nil {
		return nil, err
	}
	dao := student.ToDao()
	return &dao, nil
}

func (s *StudentServiceImpl) getStudentById(institutionId uuid.UUID, id uuid.UUID) (
	*model.Student, *httperror.HttpError,
) {
	idString := id.String()
	student, err := s.studentRepository.GetOne(dto.GetStudentFilter{ID: &idString})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentNotFound
		}
		return nil, ErrCreateStudent
	}
	if student.InstitutionID != institutionId {
		return nil, httperror.ForbiddenError
	}
	return student, nil
}

func (s *StudentServiceImpl) UpdateStudent(institutionId uuid.UUID, id uuid.UUID, dto dto.UpdateStudentDto) (
	*dao.StudentDao, *httperror.HttpError,
) {
	student, httperr := s.getStudentById(institutionId, id)
	if httperr != nil {
		return nil, httperr
	}

	if dto.FirstName != nil {
		student.FirstName = *dto.FirstName
	}

	if dto.LastName != nil {
		student.LastName = *dto.LastName
	}

	if dto.ClassID != nil {
		// get class, update class id
		student.ClassID = *dto.ClassID
	}

	err := s.studentRepository.Upsert(*student)
	if err != nil {
		return nil, ErrUpdateStudent
	}

	dao := student.ToDao()
	return &dao, nil
}

func (s *StudentServiceImpl) DeleteStudent(institutionId uuid.UUID, id uuid.UUID) *httperror.HttpError {
	student, httperr := s.getStudentById(institutionId, id)
	if httperr != nil {
		return httperr
	}

	err := s.studentRepository.Delete(*student)
	if err != nil {
		return ErrUpdateStudent
	}
	return nil
}

var ErrCreateStudent = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating student"),
}

var ErrDuplicateNisn = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Student with the same NISN already exists"),
}

var ErrDuplicateEmail = &httperror.HttpError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Student with the same email already exists"),
}

var ErrStudentNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Student not found"),
}

var ErrUpdateStudent = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating student"),
}
