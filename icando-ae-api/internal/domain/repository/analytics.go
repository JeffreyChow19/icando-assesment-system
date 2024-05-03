package repository

import (
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"time"

	"github.com/google/uuid"
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
		// join teachers only if teacherId filter is active
		// prevent COUNT on aggregating duplicate student_quiz for amount of teacher in class
		query = query.Joins("NATURAL JOIN quiz_classes").Joins("NATURAL JOIN class_teacher")
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

func (r *AnalyticsRepository) GetStudentQuizCompetency(studentID uuid.UUID) (*[]dao.GetStudentQuizCompetencyDao, error) {
	query := r.db.Table("student_quizzes sq").
		Select("c.numbering, c.name, SUM(sqc.correct_count) AS correct_sum, SUM(sqc.total_count) AS total_sum").
		Joins("INNER JOIN student_quiz_competencies sqc ON sq.id = sqc.student_quiz_id").
		Joins("INNER JOIN competencies c ON sqc.competency_id = c.id").
		Where("sq.student_id = ?", studentID).
		Group("c.numbering, c.name, sqc.competency_id")

	var results []dao.GetStudentQuizCompetencyDao
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *AnalyticsRepository) GetStudentQuizzes(studentID uuid.UUID) (*[]dao.GetStudentQuizzesDao, error) {
	query := r.db.Table("student_quizzes sq").
		Select("sq.total_score, sq.correct_count, sq.completed_at, q.name, q.passing_grade, sq.id, sq.quiz_id").
		Joins("INNER JOIN quizzes q ON sq.quiz_id = q.id").
		Where("sq.student_id = ? AND total_score IS NOT NULL", studentID).
		Order("sq.completed_at DESC")

	var results []dao.GetStudentQuizzesDao
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *AnalyticsRepository) GetTeacherDashboardOverview(id uuid.UUID) (*dao.DashboardOverviewDao, error) {
	var numClasses, numStudents, numQuizzes int

	currentTime := time.Now()
	if err := r.db.Raw(`WITH taught_classes_id AS (SELECT class_id FROM class_teacher WHERE teacher_id = ?)
	SELECT
		(SELECT COUNT(DISTINCT class_id) FROM taught_classes_id) as num_classes,
		(SELECT COUNT(DISTINCT id) FROM students WHERE class_id IN (SELECT class_id FROM taught_classes_id)) as num_students,
		(SELECT COUNT(DISTINCT quiz_id) FROM quiz_classes qc LEFT JOIN quizzes q ON qc.quiz_id = q.id WHERE q.start_at < ? AND q.end_at > ? AND qc.class_id IN (SELECT class_id FROM taught_classes_id)) as num_quizzes`, id, currentTime, currentTime).Row().Scan(&numClasses, &numStudents, &numQuizzes); err != nil {
		return nil, err
	}

	dashboardDao := dao.DashboardOverviewDao{
		TotalClass:       numClasses,
		TotalStudent:     numStudents,
		TotalOngoingQuiz: numQuizzes,
	}

	return &dashboardDao, nil
}
