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
	Teachers      []Teacher `gorm:"many2many:class_teacher;"`
	Institution   *Institution
	Students      []Student
	Quizzes       []Quiz `gorm:"many2many:quiz_classes;"`
}

func (s Class) ToDao(option dto.GetClassFilter) dao.ClassDao {
	classDao := dao.ClassDao{
		ID:            s.ID,
		Name:          s.Name,
		Grade:         s.Grade,
		InstitutionID: s.InstitutionID,
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

	if option.WithQuizRelation && s.Quizzes != nil && len(s.Quizzes) != 0 {
		quizzes := make([]dao.QuizDao, 0)

		for _, quiz := range s.Quizzes {
			quizzes = append(quizzes, quiz.ToDao(true))
		}

		classDao.Quizzes = quizzes
	}

	return classDao
}
