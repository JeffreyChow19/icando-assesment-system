package repository

import (
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type InstitutionRepository struct {
	db *gorm.DB
}

func NewInstutionRepository(db *lib.Database) InstitutionRepository {
	return InstitutionRepository{
		db: db.DB,
	}
}

// TO DO: ADD PAGINATION AND FILTER
// BUT THIS IS PROLLY NOT NEEDED
func (r *InstitutionRepository) GetAllInstitution() ([]model.Institution, error) {
	institutions := []model.Institution{}
	err := r.db.Find(&institutions).Error
	return institutions, err
}

func (r *InstitutionRepository) GetInstitution(filter dto.GetOneInstitutionFilter) (*model.Institution, error) {
	var institution model.Institution
	query := r.db
	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Slug != nil {
		query = query.Where("slug = ?", filter.Slug)
	}
	if filter.Nis != nil {
		query = query.Where("nis = ?", filter.Nis)
	}
	err := query.First(&institution).Error
	return &institution, err
}
