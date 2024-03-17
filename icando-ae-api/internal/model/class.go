package model

import (
	"icando/internal/model/dao"
	"icando/internal/model/dto"

	"github.com/google/uuid"
)

type Class struct {
	Model
	Name          string
	Grade         string
	InstitutionID uuid.UUID
	TeacherID     uuid.UUID
	Teacher       *Teacher
	Institution   *Institution
	Students      []Student
}

func (s Class) ToDao(option dto.GetClassFilter) dao.ClassDao {
	classDao := dao.ClassDao{
		ID:           s.ID,
		Name:         s.Name,
		Grade:        s.Grade,
		InstituionID: s.InstitutionID,
		TeacherID:    s.TeacherID,
	}

	if option.WithTeacherRelation && s.Teacher != nil {
		teacherDao := s.Teacher.ToDao()
		classDao.Teacher = &teacherDao
	}

	if option.WithInstitutionRelation && s.Institution != nil {
		institutionDao := s.Institution.ToDao()
		classDao.Institution = &institutionDao
	}

	if option.WithStudentRelation && s.Students != nil {
		students := make([]dao.StudentDao, 0)

		for _, student := range s.Students {
			students = append(students, student.ToDao())
		}

		classDao.Students = students
	}

	return classDao
}
