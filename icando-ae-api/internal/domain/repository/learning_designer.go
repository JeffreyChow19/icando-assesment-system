package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
)

type LearningDesignerRepository struct {
	db *gorm.DB
}

func NewLearningDesignerRepository(db *lib.Database) LearningDesignerRepository {
	return LearningDesignerRepository{
		db: db.DB,
	}
}

func (r *LearningDesignerRepository) Create(user *model.LearningDesigner) error {
	return r.db.Create(&user).Error
}

func (r *LearningDesignerRepository) FindLearningDesigner(filter dto.GetLearningDesignerFilter) (*model.LearningDesigner, error) {
	var user model.LearningDesigner

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

func (r *LearningDesignerRepository) UpdateUserInfo(user *model.LearningDesigner) error {
	return r.db.Save(&user).Error
}
