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

func (r *TeacherRepository) GetTeacher(filter dto.GetTeacherFilter) (*model.Teacher, error) {
	var user model.Teacher

	if filter.ID != nil {
		if err := r.db.Where("id = ?", filter.ID.String()).First(&user).Error; err != nil {
			return nil, err
		}
	} else if filter.Email != nil {
		if err := r.db.Where("email = ?", *filter.Email).First(&user).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Invalid filter")
	}

	return &user, nil
}

func (r *TeacherRepository) UpdateTeacher(teacher *model.Teacher) error {
	return r.db.Save(&teacher).Error
}
