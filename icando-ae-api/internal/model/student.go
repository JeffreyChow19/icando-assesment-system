package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
)

type Student struct {
	Model
	FirstName     string
	LastName      string
	Nisn          string
	Email         string
	InstitutionID uuid.UUID
	ClassID       uuid.UUID
	Institution   *Institution
	Class         *Class
}

func (s Student) ToDao() dao.StudentDao {
	return dao.StudentDao{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Nisn:      s.Nisn,
		Email:     &s.Email,
		ClassID:   s.ClassID,
	}
}
