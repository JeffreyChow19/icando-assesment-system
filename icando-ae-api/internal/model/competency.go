package model

import (
	"icando/internal/model/dao"
)

type Competency struct {
	Model
	Numbering   string
	Name        string
	Description string
}

func (c Competency) ToDao() dao.CompetencyDao {
	return dao.CompetencyDao{
		ID:          c.ID,
		Numbering:   &c.Numbering,
		Name:        &c.Name,
		Description: &c.Description,
	}
}
