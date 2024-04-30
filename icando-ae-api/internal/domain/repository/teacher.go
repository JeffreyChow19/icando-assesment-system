package repository

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"time"
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

func (r *TeacherRepository) GetTeacherDashboardOverview(id uuid.UUID) (*dao.DashboardOverviewDao, error) {
	var numClasses, numStudents, numQuizzes int

	currentTime := time.Now()
	if err := r.db.Raw(`
		SELECT 
    (SELECT COUNT(DISTINCT class_id) FROM class_teacher WHERE teacher_id = ?) as num_classes,
    (SELECT COUNT(DISTINCT id) FROM students WHERE class_id IN (SELECT class_id FROM class_teacher WHERE teacher_id = ?)) as num_students,
    (SELECT COUNT(DISTINCT quiz_id) FROM quiz_classes WHERE EXISTS (SELECT 1 FROM quizzes WHERE quizzes.id = quiz_classes.quiz_id AND quizzes.start_at < ? AND quizzes.end_at > ?) AND quiz_classes.class_id IN (SELECT class_id FROM class_teacher WHERE teacher_id = ?)) as num_quizzes`, id, id, currentTime, currentTime, id).Row().Scan(&numClasses, &numStudents, &numQuizzes); err != nil {
		return nil, err
	}

	dashboardDao := dao.DashboardOverviewDao{
		TotalClass:       numClasses,
		TotalStudent:     numStudents,
		TotalOngoingQuiz: numQuizzes,
	}

	return &dashboardDao, nil
}
