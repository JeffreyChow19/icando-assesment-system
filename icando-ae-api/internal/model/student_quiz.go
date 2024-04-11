package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"icando/internal/model/enum"
	"time"
)

type StudentQuiz struct {
	Model
	TotalScore     int
	StartedAt      *time.Time `gorm:"type:timestamptz"`
	CompletedAt    *time.Time `gorm:"type:timestamptz"`
	Status         enum.QuizStatus
	QuizID         uuid.UUID `gorm:"column:quiz_id"`
	Quiz           *Quiz
	StudentID      uuid.UUID `gorm:"column:student_id"`
	Student        *Student
	StudentAnswers []StudentAnswer
}

func (sq *StudentQuiz) ToDao() (*dao.StudentQuizDao, error) {
	studentQuizDao := dao.StudentQuizDao{
		ID:          sq.ID,
		CreatedAt:   sq.CreatedAt,
		UpdatedAt:   sq.UpdatedAt,
		TotalScore:  sq.TotalScore,
		StartedAt:   sq.StartedAt,
		CompletedAt: sq.CompletedAt,
		Status:      sq.Status,
		QuizID:      sq.QuizID,
	}

	if sq.Quiz != nil {
		quizDao := sq.Quiz.ToDao()

		studentQuizDao.Quiz = &quizDao
	}

	if sq.StudentAnswers != nil {
		answersDao := make([]dao.StudentAnswerDao, 0)

		for _, answer := range sq.StudentAnswers {
			answerDao, err := answer.ToDao()

			if err != nil {
				return nil, err
			}

			answersDao = append(answersDao, *answerDao)
		}

		studentQuizDao.StudentAnswers = answersDao
	}

	if sq.Student != nil {
		studentDao := sq.Student.ToDao()
		studentQuizDao.Student = &studentDao
	}

	return &studentQuizDao, nil
}
