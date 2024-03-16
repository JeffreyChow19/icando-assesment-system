package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
)

type Class struct {
	Model
	Name          string
	Grade         string
	InstitutionID uuid.UUID
	Teachers      []Teacher `gorm:"many2many:class_teacher;"`
	Institution   *Institution
	Students      []Student
}

func (s Class) ToDao(option dto.GetClassFitler) dao.ClassDao {
	classDao := dao.ClassDao{
		ID:           s.ID,
		Name:         s.Name,
		Grade:        s.Grade,
		InstituionID: s.InstitutionID,
	}

	if option.WithTeacherRelation {
		teachersDao := make([]dao.TeacherDao, 0)

		for _, teacher := range s.Teachers {
			teachersDao = append(teachersDao, teacher.ToDao())
		}

		classDao.Teachers = teachersDao
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
