package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
)

type Institution struct {
	ID   uuid.UUID `gorm:"primarykey"`
	Name string
	Nis  string
	Slug string
}

func (s Institution) ToDao() dao.InstitutionDao {
	return dao.InstitutionDao{
		ID:   s.ID,
		Name: s.Name,
		Nis:  s.Nis,
		Slug: s.Slug,
	}
}
