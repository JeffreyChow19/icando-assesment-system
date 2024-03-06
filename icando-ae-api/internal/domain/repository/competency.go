package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"math"
)

type CompetencyRepository struct {
	db *gorm.DB
}

func NewCompetencyRepository(db *lib.Database) CompetencyRepository {
	return CompetencyRepository{
		db: db.DB,
	}
}

var ErrCompetencyNotFound = errors.New("competency not found")

func (r *CompetencyRepository) GetCompetencyById(id uuid.UUID) (*model.Competency, error) {
	var competency model.Competency
	if err := r.db.Where("id = ?", id).First(&competency).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCompetencyNotFound
		}
		return nil, err
	}
	return &competency, nil
}

func (r *CompetencyRepository) GetCompetencyByNumbering(numbering string) (*model.Competency, error) {
	var competency model.Competency
	if err := r.db.Where("numbering = ?", numbering).First(&competency).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &competency, nil
}

func (r *CompetencyRepository) GetAllCompetencies(filter dto.GetAllCompetenciesFilter) ([]model.Competency, *dao.MetaDao, error) {
	query := r.db.Model(&model.Competency{})

	if filter.Numbering != nil {
		query.Where("numbering ILIKE ?", fmt.Sprintf("%s%", filter.Numbering))
	}

	if filter.Name != nil {
		query.Where("name ILIKE ?", fmt.Sprintf("%s%", filter.Name))
	}

	var totalItem int64
	err := query.Session(&gorm.Session{}).Count(&totalItem).Error
	if err != nil {
		return nil, nil, err
	}

	meta := dao.MetaDao{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(float64(totalItem) / float64(filter.Limit))),
	}

	Paginate(query, filter.Page, filter.Limit)

	var competencies []model.Competency
	err = query.Session(&gorm.Session{}).Find(&competencies).Error

	return competencies, &meta, nil
}

func (r *CompetencyRepository) CreateCompetency(competency model.Competency) error {
	return r.db.Create(&competency).Error
}

func (r *CompetencyRepository) UpdateCompetency(competency model.Competency) error {
	return r.db.Save(&competency).Error
}

func (r *CompetencyRepository) DeleteCompetency(competency model.Competency) error {
	return r.db.Delete(&competency).Error
}
