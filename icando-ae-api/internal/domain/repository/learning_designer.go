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

func (r *LearningDesignerRepository) Create(user *model.LearningDesigner) error {
	return r.db.Create(&user).Error
}

func (r *LearningDesignerRepository) FindUserById(id uuid.UUID) (*model.LearningDesigner, error) {
	var user model.LearningDesigner
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) FindUserByEmail(email string) (*model.LearningDesigner, error) {
	var user model.LearningDesigner
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) FindUserByName(firstName string, lastName string) (*model.LearningDesigner, error) {
	var user model.LearningDesigner
	err := r.db.Where("first_name = ? AND last_name = ?", firstName, lastName).First(&user).Error
	return &user, err
}

func (r *LearningDesignerRepository) UpdateUserInfo(user *model.LearningDesigner) error {
	return r.db.Save(&user).Error
}