package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"icando/internal/model/enum"
	"time"
)

type StudentQuiz struct {
	Model
	TotalScore     *float32
	CorrectCount   *int
	StartedAt      *time.Time `gorm:"type:timestamptz"`
	CompletedAt    *time.Time `gorm:"type:timestamptz"`
	Status         enum.QuizStatus
	QuizID         uuid.UUID `gorm:"column:quiz_id"`
	Quiz           *Quiz
	StudentID      uuid.UUID `gorm:"column:student_id"`
	Student        *Student
	StudentAnswers []StudentAnswer `gorm:"foreignKey:student_quiz_id"`
}

func (sq *StudentQuiz) ToDao(withQuestionAnswer bool) (*dao.StudentQuizDao, error) {
	studentQuizDao := dao.StudentQuizDao{
		ID:           sq.ID,
		CreatedAt:    sq.CreatedAt,
		UpdatedAt:    sq.UpdatedAt,
		TotalScore:   sq.TotalScore,
		CorrectCount: sq.CorrectCount,
		StartedAt:    sq.StartedAt,
		CompletedAt:  sq.CompletedAt,
		Status:       sq.Status,
		QuizID:       sq.QuizID,
		StudentID:    sq.StudentID,
	}

	if sq.Quiz != nil {
		quizDao := sq.Quiz.ToDao(withQuestionAnswer)

		studentQuizDao.Quiz = &quizDao
	}

	if sq.StudentAnswers != nil {
		answersDao := make([]dao.StudentAnswerDao, 0)

		for _, answer := range sq.StudentAnswers {
			answerDao, err := answer.ToDao(withQuestionAnswer)

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
