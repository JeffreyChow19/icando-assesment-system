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

type StudentService interface {
	AddStudent(institutionId uuid.UUID, studentDto dto.CreateStudentDto) (*dao.StudentDao, *httperror.HttpError)
}

type StudentServiceImpl struct {
	studentRepository repository.StudentRepository
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		studentRepository: studentRepository,
	}
}

func (s *StudentServiceImpl) AddStudent(institutionId uuid.UUID, studentDto dto.CreateStudentDto) (
	*dao.StudentDao, *httperror.HttpError,
) {
	// to do get class by id
	// check if class s institution == institution id
	student := model.Student{
		FirstName:     studentDto.FirstName,
		LastName:      studentDto.LastName,
		Email:         studentDto.Email,
		Nisn:          studentDto.Nisn,
		ClassID:       studentDto.ClassID,
		InstitutionID: institutionId,
	}
	err := s.studentRepository.Create(student)
	if err != nil {
		return nil, ErrCreateStudent
	}
	dao := student.ToDao()
	return &dao, nil
}

var ErrCreateStudent = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating student"),
}
