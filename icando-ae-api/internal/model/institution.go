package model

import (
	"icando/internal/model/dao"
)

type Institution struct {
	Model
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
