package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *lib.Database) TeacherRepository {
	return TeacherRepository{
		db: db.DB,
	}
}

func (r *TeacherRepository) GetAllTeacher(filter dto.GetTeacherFilter) ([]model.Teacher, error) {
	query := r.db.Table("teachers")
	result := make([]model.Teacher, 0)

	if filter.ID != nil {
		query.Where("id = ?", filter.ID.String())
	}
	if filter.Email != nil {
		query.Where("email = ?", *filter.Email)
	}
	if filter.InstitutionID != nil {
		query.Where("institution_id = ?", filter.InstitutionID.String())
	}

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TeacherRepository) GetTeacher(filter dto.GetTeacherFilter) (*model.Teacher, error) {
	var user model.Teacher

	query := r.db

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID.String())
	} else if filter.Email != nil {
		query = query.Where("email = ?", *filter.Email)
	} else {
		return nil, errors.New("Invalid filter")
	}

	if filter.WithClasses != nil && *filter.WithClasses {
		query = query.Preload("Classes")
	}

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TeacherRepository) UpdateTeacher(teacher *model.Teacher) error {
	return r.db.Save(&teacher).Error
}
