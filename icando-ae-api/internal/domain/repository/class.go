package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
	"icando/utils"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *lib.Database) ClassRepository {
	return ClassRepository{
		db: db.DB,
	}
}

func (r *ClassRepository) CreateClass(dto dto.ClassDto) (*model.Class, error) {
	class := model.Class{
		Name:         dto.Name,
		Grade:        dto.Grade,
		InstituionID: dto.InstitutionID,
		TeacherID:    dto.TeacherID,
	}

	if err := r.db.Create(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepository) UpdateClass(id uuid.UUID, dto dto.ClassDto) (*model.Class, error) {
	class := model.Class{
		ID:           id,
		Name:         dto.Name,
		Grade:        dto.Grade,
		InstituionID: dto.InstitutionID,
		TeacherID:    dto.TeacherID,
	}

	if err := r.db.Save(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepository) DeleteClass(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id.String()).Delete(&model.Class{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *ClassRepository) GetAllClass(filter dto.GetAllClassFilter) ([]model.Class, error) {
	var scopes []func(*gorm.DB) *gorm.DB
	result := make([]model.Class, 0)

	if filter.SortBy != nil && *filter.SortBy != "" {
		scopes = append(scopes, utils.QuerySortBy(*filter.SortBy, !filter.Desc))
	}

	if filter.TeacherID != nil {
		scopes = append(scopes, classWhereTeacherID(*filter.TeacherID))
	}

	if filter.InstitutionID != nil {
		scopes = append(scopes, classWhereInstitutionID(*filter.InstitutionID))
	}

	if err := r.db.Model(&model.Class{}).Scopes(scopes...).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ClassRepository) GetClass(id uuid.UUID, filter dto.GetClassFitler) (*model.Class, error) {
	var class model.Class

	query := r.db.Session(&gorm.Session{})

	if filter.WithInstitutionRelation {
		query = query.Preload("Institution")
	}

	if filter.WithStudentRelation {
		query = query.Preload("Students")
	}

	if filter.WithTeacherRelation {
		query = query.Preload("Teacher")
	}

	if err := query.Where("id = ?", id.String()).First(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func classWhereTeacherID(id uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("teached_id = ?", id.String())
	}
}

func classWhereInstitutionID(id uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("institution_id = ?", id.String())
	}
}
