package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (r *CompetencyRepository) GetOneCompetency(filter dto.GetOneCompetencyFilter) (*model.Competency, error) {
	query := r.db.Model(&model.Competency{})

	if filter.Id != uuid.Nil {
		query.Where("id = ?", filter.Id)
	}

	if filter.Numbering != nil {
		query.Where("numbering ILIKE ?", fmt.Sprintf("%s%%", *filter.Numbering))
	}

	var competency model.Competency
	err := query.First(&competency).Error
	if err != nil {
		return nil, err
	}

	return &competency, nil
}

func (r *CompetencyRepository) GetAllCompetencies(filter dto.GetAllCompetenciesFilter) (
	[]model.Competency, *dao.MetaDao, error,
) {
	query := r.db.Model(&model.Competency{})

	if filter.Numbering != nil {
		query.Where("numbering ILIKE ?", fmt.Sprintf("%s%%", *filter.Numbering))
	}

	if filter.Name != nil {
		query.Where("name ILIKE ?", fmt.Sprintf("%s%%", *filter.Name))
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

	Sort(query, true, "numbering")
	Paginate(query, filter.Page, filter.Limit)

	var competencies []model.Competency
	err = query.Session(&gorm.Session{}).Find(&competencies).Error

	return competencies, &meta, nil
}

func (r *CompetencyRepository) GetCompetenciesByIDs(competencyIDs []uuid.UUID) ([]model.Competency, error) {
	var competencies []model.Competency

	// Convert the slice of UUIDs to a slice of interfaces for the Where clause
	ids := make([]interface{}, len(competencyIDs))
	for i, id := range competencyIDs {
		ids[i] = id
	}

	// Find all Competency records where the ID is in the provided list
	err := r.db.Model(&model.Competency{}).Where("id IN (?)", ids).Find(&competencies).Error
	if err != nil {
		return nil, err
	}

	return competencies, nil
}

func (r *CompetencyRepository) CreateCompetency(competency model.Competency) (model.Competency, error) {
	err := r.db.Create(&competency).Error
	return competency, err
}

func (r *CompetencyRepository) UpdateCompetency(competency model.Competency) error {
	return r.db.Save(&competency).Error
}

func (r *CompetencyRepository) DeleteCompetency(competency model.Competency) error {
	return r.db.Delete(&competency).Error
}

func (r *CompetencyRepository) GetStudentCompetency(filter dto.GetStudentCompetencyFilter) ([]dao.StudentCompetencyDao, error) {
	if filter.StudentID == nil && filter.StudentQuizID == nil {
		return nil, errors.New("Both student id and student quiz id cannot be null")
	} else if filter.StudentID != nil && filter.StudentQuizID != nil {
		return nil, errors.New("Either student id and student quiz id should have a value")
	}

	var result []dao.StudentCompetencyDao

	if filter.StudentID != nil {
		if err := r.db.Raw(`
		SELECT
			id as competency_id,
			name as competency_name,
			coalesce(sum(total_count), 0) as total_count,
			coalesce(sum(correct_count), 0) as correct_count
		FROM competencies
		LEFT JOIN student_quiz_competencies
		ON competencies.id = student_quiz_competencies.competency_id AND student_id = ?
		GROUP BY competencies.id
	`, filter.StudentID.String()).Scan(&result).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Raw(`
		SELECT
			id as competency_id,
			name as competency_name,
			coalesce(sum(total_count), 0) as total_count,
			coalesce(sum(correct_count), 0) as correct_count
		FROM competencies
		LEFT JOIN student_quiz_competencies
		ON competencies.id = student_quiz_competencies.competency_id AND student_quiz_id = ?
		GROUP BY competencies.id
	`, filter.StudentQuizID.String()).Scan(&result).Error; err != nil {
			return nil, err
		}
	}

	if filter.StudentID != nil {
		return result, nil
	} else {
		// hacky way but more efficient <3
		// this is case for showing competencies for student quiz
		// we don't want to show competencies that does not show in the quiz
		// to check if the competency is in the quiz, we assume that if totalCount is 0
		// then the competency does not show up in the quiz
		filteredResult := make([]dao.StudentCompetencyDao, 0)

		for _, each := range result {
			if each.TotalCount > 0 {
				filteredResult = append(filteredResult, each)
			}
		}

		return filteredResult, nil
	}
}
