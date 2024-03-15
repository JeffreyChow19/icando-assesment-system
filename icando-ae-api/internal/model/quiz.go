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
	UpdatedBy    *uuid.UUID `gorm:"type:uuid"`
	PublishedAt  *time.Time `gorm:"type:timestamptz"`
	Deadline     *time.Time `gorm:"type:timestamptz"`
}

func (q Quiz) ToDao() dao.QuizDao {
	daoQuiz := dao.QuizDao{
		ID:           q.ID,
		Name:         q.Name,
		Subject:      q.Subject,
		PassingGrade: q.PassingGrade,
	}

	if q.Deadline != nil {
		daoQuiz.Deadline = q.Deadline
	}

	return daoQuiz
}
