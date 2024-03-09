package dto

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
)

type GetOneCompetencyFilter struct {
	Id        uuid.UUID
	Numbering *string
}

type GetAllCompetenciesFilter struct {
	Numbering *string `form:"numbering"`
	Name      *string `form:"name"`
	Page      int     `form:"page"`
	Limit     int     `form:"limit"`
}

type GetAllCompetenciesResponse struct {
	Competencies []dao.CompetencyDao `json:"competencies"`
	Meta         dao.MetaDao         `json:"meta"`
}

type CreateCompetencyDto struct {
	Numbering   string `json:"numbering" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCompetencyDto struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	Numbering   *string   `json:"numbering,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
}
