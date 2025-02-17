package model

import (
	"icando/internal/model/base"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	Model
	Name            *string
	Subject         base.StringArray `gorm:"type:text[]"`
	PassingGrade    float64
	ParentQuiz      *uuid.UUID
	CreatedBy       uuid.UUID  `gorm:"type:uuid;not null"`
	Creator         *Teacher   `gorm:"foreignKey:CreatedBy"`
	UpdatedBy       *uuid.UUID `gorm:"type:uuid"`
	Updater         *Teacher   `gorm:"foreignKey:UpdatedBy"`
	PublishedAt     *time.Time `gorm:"type:timestamptz"`
	Duration        *int
	StartAt         *time.Time `gorm:"type:timestamptz"`
	EndAt           *time.Time `gorm:"type:timestamptz"`
	Questions       []Question
	Classes         []Class `gorm:"many2many:quiz_classes;"`
	HasNewerVersion *bool   `gorm:"-"`
}

type QuizClass struct {
	QuizID  uuid.UUID
	ClassID uuid.UUID
}

func (q Quiz) ToDao(withAnswer bool) dao.QuizDao {
	daoQuiz := dao.QuizDao{
		ID:           q.ID,
		Name:         q.Name,
		Subject:      q.Subject,
		PassingGrade: q.PassingGrade,
		PublishedAt:  q.PublishedAt,
		Duration:     q.Duration,
		StartAt:      q.StartAt,
		EndAt:        q.EndAt,
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
			questionDao, err := question.ToDao(withAnswer)
			if err != nil {
				continue
			}

			questions = append(questions, *questionDao)
		}
		daoQuiz.Questions = questions
	}

	if q.Classes != nil && len(q.Classes) != 0 {
		classes := make([]dao.ClassDao, 0)

		for _, class := range q.Classes {
			classes = append(classes, class.ToDao(dto.GetClassFilter{}))
		}

		daoQuiz.Classes = classes
	}

	if q.HasNewerVersion != nil {
		daoQuiz.HasNewerVersion = q.HasNewerVersion
	}

	return daoQuiz
}
