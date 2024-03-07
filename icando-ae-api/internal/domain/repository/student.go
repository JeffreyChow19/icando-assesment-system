package repository

import (
	"fmt"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"math"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *lib.Database) StudentRepository {
	return StudentRepository{
		db: db.DB,
	}
}

func (r *StudentRepository) GetAllStudent(filter dto.GetAllStudentsFilter) ([]model.Student, *dao.MetaDao, error) {
	query := r.db
	if filter.IncludeInstitution {
		query.Preload("Institution")
	}
	if filter.IncludeClass {
		query.Preload("Class")
	}
	if filter.InstitutionID != nil {
		query.Where("institution_id = ?", filter.InstitutionID)
	}
	if filter.Name != nil {
		query.Where("concat(first_name, ' ', last_name) ilike ?", fmt.Sprintf("%s%", filter.Name))
	}
	if filter.ClassID != nil {
		query.Where("class_id = ?", filter.ClassID)
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

	var students []model.Student
	err = query.Session(&gorm.Session{}).Find(&students).Error

	return students, &meta, err
}

func (r *StudentRepository) GetOne(filter dto.GetStudentFilter) (*model.Student, error) {
	query := r.db
	if filter.IncludeInstitution {
		query.Preload("Institution")
	}
	if filter.IncludeClass {
		query.Preload("Class")
	}
	if filter.Nisn != nil {
		query.Where("nisn = ?", filter.Nisn)
	}
	if filter.ID != nil {
		query.Where("id = ?", filter.ID)
	}

	var student model.Student
	err := query.Session(&gorm.Session{}).First(&student).Error
	return &student, err
}

func (r *StudentRepository) Create(student model.Student) error {
	return r.db.Create(&student).Error
}

func (r *StudentRepository) Upsert(student model.Student) error {
	return r.db.Save(&student).Error
}

func (r *StudentRepository) Delete(student model.Student) error {
	return r.db.Delete(&student).Error
}
