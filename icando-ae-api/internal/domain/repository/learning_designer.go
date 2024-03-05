package repository

import (
	"github.com/google/uuid"
	"icando/internal/model"
	"icando/lib"
	"gorm.io/gorm"
)

type LearningDesignerRepository struct {
	db *gorm.DB
}

func NewLearningDesignerRepository(db *lib.Database) LearningDesignerRepository {
	return LearningDesignerRepository{
		db: db.DB,
	}
}

func (r *LearningDesignerRepository) FindUserById(id uuid.UUID) (*model.LearningDesiner, error) {
	var user model.LearningDesiner
	err := r.db.Preload("Learning_Designer").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) FindUserByEmail(email string) (*model.LearningDesiner, error) {
	var user model.LearningDesiner
	err := r.db.Preload("Learning_Designer").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) FindUserByName(firstName string, lastName string) (*model.LearningDesiner, error) {
	var user model.LearningDesiner
	err := r.db.Preload("Learning_Designer").Where("first_name = ? AND last_name = ?", firstName, lastName).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) UpdateUserInfo(user *model.LearningDesiner) error {
	return r.db.Save(&user).Error
}