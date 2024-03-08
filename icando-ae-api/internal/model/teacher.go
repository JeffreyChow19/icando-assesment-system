package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
)

type Teacher struct {
	ID            uuid.UUID `gorm:"primarykey"`
	FirstName     string
	LastName      string
	Email         string
	Password      string
	InstitutionID uuid.UUID
	Institution   *Institution
}

func (s Teacher) ToDao() dao.TeacherDao {
	return dao.TeacherDao{
		ID:            s.ID,
		FirstName:     s.FirstName,
		LastName:      s.LastName,
		Email:         &s.Email,
		InstitutionID: s.InstitutionID,
	}
}
