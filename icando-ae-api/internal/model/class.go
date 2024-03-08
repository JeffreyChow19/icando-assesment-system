package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
)

type Class struct {
	ID           uuid.UUID `gorm:"primarykey"`
	Name         string
	Grade        string
	InstituionID uuid.UUID
	TeacherID    uuid.UUID
	Teacher      *Teacher
	Institution  *Institution
	Students     []Student
}

func (s Class) ToDao(option dto.GetClassFitler) dao.ClassDao {
	classDao := dao.ClassDao{
		ID:           s.ID,
		Name:         s.Name,
		Grade:        s.Grade,
		InstituionID: s.InstituionID,
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
