package repository

import (
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"

	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *lib.Database) AnalyticsRepository {
	return AnalyticsRepository{
		db: db.DB,
	}
}

func (r *AnalyticsRepository) GetQuizPerformance(filter *dto.GetQuizPerformanceFilter) (*dao.QuizPerformanceDao, error) {
	query := r.db.Table("student_quizzes").
		Joins("JOIN quizzes ON student_quizzes.quiz_id = quizzes.id").
		Joins("NATURAL JOIN quiz_classes").
		Joins("NATURAL JOIN class_teacher").
		Select(`
			COUNT(CASE WHEN total_score >= passing_grade THEN 1 END) AS quizzes_passed, 
			COUNT(CASE WHEN total_score < passing_grade THEN 1 END) AS quizzes_failed
			`)

	if filter.QuizID != nil {
		query = query.Where("quiz_id = ?", filter.QuizID)
	}

	if filter.StudentID != nil {
		query = query.Where("student_id = ?", filter.StudentID)
	}

	if filter.TeacherID != nil {
		query = query.Where("teacher_id = ?", filter.TeacherID)
	}

	var result dao.QuizPerformanceDao
	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AnalyticsRepository) GetLatestSubmissions(filter *dto.GetLatestSubmissionsFilter) (*[]dao.GetLatestSubmissionsDao, error) {
	query := r.db.Table("student_quizzes").
		Select("classes.name as class_name, classes.grade, quizzes.name as quiz_name, students.first_name, students.last_name, student_quizzes.completed_at").
		Joins("JOIN quizzes ON student_quizzes.quiz_id = quizzes.id").
		Joins("JOIN quiz_classes ON student_quizzes.quiz_id = quiz_classes.quiz_id").
		Joins("JOIN classes ON quiz_classes.class_id = classes.id").
		Joins("JOIN class_teacher ON classes.id = class_teacher.class_id").
		Joins("JOIN students ON student_quizzes.student_id = students.id").
		Where("student_quizzes.status = ?", "SUBMITTED")

	if filter.TeacherID != nil {
		query = query.Where("class_teacher.teacher_id = ?", filter.TeacherID)
	}

	query = query.Order("student_quizzes.completed_at desc").Limit(10)

	var results []dao.GetLatestSubmissionsDao
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return &results, nil
}
