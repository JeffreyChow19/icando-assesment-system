package repository

import (
	"fmt"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/lib"
	"icando/utils"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	tx := r.db.Begin()

	class := model.Class{
		Name:          dto.Name,
		Grade:         dto.Grade,
		InstitutionID: dto.InstitutionID,
	}

	if err := tx.Create(&class).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var values []string

	for _, teacherId := range dto.TeacherIDs {
		values = append(values, fmt.Sprintf("('%s', '%s')", class.ID.String(), teacherId.String()))
	}

	if values != nil && len(values) != 0 {
		query := fmt.Sprintf("INSERT INTO class_teacher(class_id, teacher_id) VALUES %s", strings.Join(values, ", "))

		if err := tx.Exec(query).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepository) UpdateClass(id uuid.UUID, classDto dto.ClassDto) (*model.Class, error) {

	oldClass, err := r.GetClass(id, dto.GetClassFilter{WithTeacherRelation: true})

	if err != nil {
		return nil, err
	}

	tx := r.db.Begin()

	class := model.Class{
		Name:          classDto.Name,
		Grade:         classDto.Grade,
		InstitutionID: classDto.InstitutionID,
	}
	class.ID = id

	if err := tx.Save(&class).Error; err != nil {
		return nil, err
	}

	// to delete teacher
	toDeleteIds := []string{}

	for _, teacher := range oldClass.Teachers {
		shouldDelete := true

		for _, newTeacherId := range classDto.TeacherIDs {
			if newTeacherId.String() == teacher.ID.String() {
				shouldDelete = false
				break
			}
		}

		if shouldDelete {
			toDeleteIds = append(toDeleteIds, teacher.ID.String())
		}
	}

	if toDeleteIds != nil && len(toDeleteIds) != 0 {
		if err := tx.Exec("DELETE FROM class_teacher WHERE class_id = ? AND teacher_id IN ?", oldClass.ID.String(), toDeleteIds).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var toAddIds []string

	for _, teacherId := range classDto.TeacherIDs {
		shouldAdd := true

		for _, oldTeacher := range oldClass.Teachers {
			if oldTeacher.ID.String() == teacherId.String() {
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			toAddIds = append(toAddIds, fmt.Sprintf("('%s', '%s')", oldClass.ID.String(), teacherId.String()))
		}
	}

	if toAddIds != nil && len(toAddIds) != 0 {
		query := fmt.Sprintf("INSERT INTO class_teacher(class_id, teacher_id) VALUES %s", strings.Join(toAddIds, ", "))

		if err := tx.Exec(query).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
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
	query := r.db
	result := make([]model.Class, 0)

	if filter.SortBy != nil && *filter.SortBy != "" {
		utils.QuerySortBy(query, *filter.SortBy, !filter.Desc)
	}

	if filter.TeacherID != nil {
		var classIds []ClassIds
		ids := make([]string, 0)

		if err := r.db.Raw("SELECT class_id FROM class_teacher WHERE teacher_id = ?", filter.TeacherID.String()).Scan(&classIds).Error; err != nil {
			return nil, err
		}

		for _, classId := range classIds {
			ids = append(ids, classId.ClassID.String())
		}

		query.Where("id IN ?", ids)
	}

	if filter.InstitutionID != nil {
		query.Where("institution_id = ?", filter.InstitutionID.String())
	}

	query = query.Preload("Teachers")

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ClassRepository) GetClass(id uuid.UUID, filter dto.GetClassFilter) (*model.Class, error) {
	var class model.Class

	query := r.db.Session(&gorm.Session{})

	if filter.WithInstitutionRelation {
		query = query.Preload("Institution")
	}

	if filter.WithStudentRelation {
		query = query.Preload("Students")
	}

	if filter.WithTeacherRelation {
		query = query.Preload("Teachers")
	}

	if filter.WithQuizRelation {
		query = query.Preload("Quizzes")
	}

	if err := query.Where("id = ?", id.String()).First(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

type ClassIds struct {
	ClassID uuid.UUID
}
