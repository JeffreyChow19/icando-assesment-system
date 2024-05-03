package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
)

type Teacher struct {
	Model
	FirstName     string
	LastName      string
	Email         string
	Password      string
	InstitutionID uuid.UUID
	Institution   *Institution
	Classes       []Class `gorm:"many2many:class_teacher;"`
	Role          enum.TeacherRole
}

func (s Teacher) IsTeachingClass(classID uuid.UUID) (bool, error) {
	if s.Classes == nil {
		return false, errors.New("Class data is not loader")
	}

	isFound := false

	for _, class := range s.Classes {
		if class.ID.String() == classID.String() {
			isFound = true
			break
		}
	}

	return isFound, nil
}

func (s Teacher) IsTeachingClasses(classIds []uuid.UUID) (bool, error) {
	for _, id := range classIds {
		isFound, err := s.IsTeachingClass(id)
		if err != nil {
			return false, err
		}
		if isFound {
			return true, nil
		}
	}
	return false, nil
}

func (s Teacher) ToDao() dao.TeacherDao {
	teacherDao := dao.TeacherDao{
		ID:            s.ID,
		FirstName:     s.FirstName,
		LastName:      s.LastName,
		Email:         &s.Email,
		InstitutionID: s.InstitutionID,
	}

	if s.Classes != nil {
		classesDao := make([]dao.ClassDao, 0)

		for _, class := range s.Classes {
			classesDao = append(classesDao, class.ToDao(dto.GetClassFilter{}))
		}

		teacherDao.Classes = classesDao
	}

	return teacherDao
}
