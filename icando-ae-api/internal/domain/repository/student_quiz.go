package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"math"
)

type StudentQuizRepository struct {
	db             *gorm.DB
	quizRepository QuizRepository
}

func NewStudentQuizRepository(db *lib.Database, quizRepository QuizRepository) StudentQuizRepository {
	return StudentQuizRepository{
		db:             db.DB,
		quizRepository: quizRepository,
	}
}

func (r *StudentQuizRepository) GetStudentQuiz(filter dto.GetStudentQuizFilter) (*model.StudentQuiz, error) {
	query := r.db.Session(&gorm.Session{})

	if filter.WithAnswers {
		query = query.Preload("StudentAnswers")
	}

	if filter.WithStudent {
		query = query.Preload("Student")
	}

	query = query.Where("id = ?", filter.ID.String())

	var studentQuiz model.StudentQuiz

	if err := query.First(&studentQuiz).Error; err != nil {
		return nil, err
	}

	if filter.WithQuizOverview || filter.WithQuizQuestions {
		quiz, err := r.quizRepository.GetQuiz(
			dto.GetQuizFilter{
				ID:            studentQuiz.QuizID,
				WithQuestions: filter.WithQuizQuestions,
			},
		)

		if err != nil {
			return nil, err
		}

		studentQuiz.Quiz = quiz
	}

	return &studentQuiz, nil
}

func (r *StudentQuizRepository) CreateStudentQuiz(studentQuiz model.StudentQuiz) (model.StudentQuiz, error) {
	err := r.db.Create(&studentQuiz).Error

	return studentQuiz, err
}

func (r *StudentQuizRepository) UpdateStudentQuiz(studentQuiz model.StudentQuiz) error {
	return r.db.Save(&studentQuiz).Error
}

func (r *StudentQuizRepository) UpdateAnswer(answer model.StudentAnswer) error {
	return r.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(&answer).Error
}

func (r *StudentQuizRepository) GetStudentQuizzes(filter dto.GetStudentQuizzesFilter) (
	[]dao.GetStudentQuizzesDao,
	*dao.MetaDao,
	error,
) {
	baseQuery := `
	WITH config AS (
	  SELECT 
          ?::uuid AS quiz_id,
	      ?::uuid AS class_id,
	      ?::varchar AS student_name,
	  	  ?::quiz_status AS quiz_status,
          ?::uuid AS teacher_id
	),
	quiz AS (
	    SELECT q.id id FROM quizzes q LEFT JOIN config c ON TRUE
		WHERE
			CASE WHEN c.quiz_id IS NOT NULL THEN q.id = c.quiz_id
			ELSE TRUE
		END
	)
	`

	tableQuery := `
	FROM quiz 
		LEFT JOIN config ON true
		INNER JOIN student_quizzes sq ON quiz.id = sq.quiz_id 
		INNER JOIN students s ON s.id = sq.student_id 
		INNER JOIN classes c ON c.id = s.class_id
		INNER JOIN class_teacher ct ON ct.class_id = c.id
	WHERE
	    CASE WHEN config.class_id IS NOT NULL 
	        THEN c.id = config.class_id 
	        ELSE TRUE
		END
		AND CASE WHEN config.student_name IS NOT NULL
			THEN '%' || config.student_name || '%' ILIKE CONCAT(s.first_name, ' ', s.last_name)
			ELSE TRUE
		END
		AND CASE WHEN config.quiz_status IS NOT NULL
			THEN sq.status = config.quiz_status
			ELSE TRUE
		END
		AND CASE WHEN config.teacher_id IS NOT NULL
			THEN ct.teacher_id = config.teacher_id
			ELSE TRUE
		END
`

	countQuery := baseQuery + `SELECT COUNT(sq.id)` + tableQuery
	query := baseQuery + `
	SELECT 
	    concat(s.first_name, ' ', s.last_name) name,
	    sq.correct_count correct_count,
	    sq.total_score total_score,
	    sq.status status,
	    sq.completed_at completed_at
	` + tableQuery + ` ORDER BY sq.completed_at DESC`

	var totalItem int64
	results := []dao.GetStudentQuizzesDao{}

	err := r.db.Raw(
		countQuery, filter.QuizID, filter.ClassID, filter.StudentName, filter.QuizStatus, filter.TeacherID,
	).Scan(&totalItem).Error
	if err != nil {
		return results, nil, err
	}
	meta := dao.MetaDao{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(float64(totalItem) / float64(filter.Limit))),
	}

	searchQuery := r.db.Raw(
		query, filter.QuizID, filter.ClassID, filter.StudentName, filter.QuizStatus, filter.TeacherID,
	)
	Paginate(searchQuery, filter.Page, filter.Limit)
	err = searchQuery.Scan(&results).Error
	if err != nil {
		return results, nil, nil
	}

	return results, &meta, nil
}
