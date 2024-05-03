package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
)

type Student struct {
	Model
	FirstName     string
	LastName      string
	Nisn          string
	Email         string
	InstitutionID uuid.UUID    `gorm:"type:uuid;not null"`
	ClassID       *uuid.UUID   `gorm:"type:uuid"`
	Institution   *Institution `gorm:"foreignKey:InstitutionID"`
	Class         *Class       `gorm:"foreignKey:ClassID"`
}

func (s Student) ToDao() dao.StudentDao {
	studentDao := dao.StudentDao{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Nisn:      s.Nisn,
		Email:     s.Email,
		ClassID:   s.ClassID,
	}

	if s.Class != nil {
		classDao := s.Class.ToDao(dto.GetClassFilter{})
		studentDao.Class = &classDao
	}

	return studentDao
}
