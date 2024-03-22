package model

import (
	"github.com/google/uuid"
	"icando/internal/model/dao"
	"time"
)

type Quiz struct {
	Model
	Name         *string
	Subject      *string
	PassingGrade float64
	ParentQuiz   *uuid.UUID
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null"`
	Creator			 *Teacher		`gorm:"foreignKey:CreatedBy"`
	UpdatedBy    *uuid.UUID `gorm:"type:uuid"`
	Updater			 *Teacher		`gorm:"foreignKey:UpdatedBy"`
	PublishedAt  *time.Time `gorm:"type:timestamptz"`
	Deadline     *time.Time `gorm:"type:timestamptz"`
	Questions		 []Question	
}

func (q Quiz) ToDao() dao.QuizDao {
	daoQuiz := dao.QuizDao{
		ID:           q.ID,
		Name:         q.Name,
		Subject:      q.Subject,
		PassingGrade: q.PassingGrade,
		PublishedAt:  q.PublishedAt,
		Deadline:     q.Deadline,
	}

	if q.Creator != nil {
		creatorDao := q.Creator.ToDao()
		daoQuiz.Creator = &creatorDao
	}

	if q.Updater != nil {
		updaterDao := q.Updater.ToDao()
		daoQuiz.Updater = &updaterDao
	}

	if q.Questions != nil {
		questions := make([]dao.QuestionDao, 0)

		for _, question := range q.Questions {
			questionDao, err := question.ToDao()
			if err != nil {
				continue
			}

			questions = append(questions, *questionDao)
		}

		daoQuiz.Questions = questions
	}

	return daoQuiz
}
