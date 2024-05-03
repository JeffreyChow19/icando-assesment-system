package dao

import (
	"icando/internal/model/enum"
	"time"
)

type QuizPerformanceDao struct {
	QuizzesPassed int `json:"quizzesPassed"`
	QuizzesFailed int `json:"quizzesFailed"`
}

type GetLatestSubmissionsDao struct {
	ClassName   string    `json:"class_name"`
	Grade       string    `json:"grade"`
	QuizName    string    `json:"quiz_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CompletedAt time.Time `json:"completed_at"`
}

type GetStudentQuizCompetencyDao struct {
	Numbering  string `json:"numbering"`
	Name       string `json:"name"`
	CorrectSum int    `json:"correct_sum"`
	TotalSum   int    `json:"total_sum"`
}

type GetStudentQuizzesDao struct {
	TotalScore   float32          `json:"total_score" gorm:"column:total_score"`
	CorrectCount int              `json:"correct_count" gorm:"column:correct_count"`
	CompletedAt  time.Time        `json:"completed_at" gorm:"column:completed_at"`
	Name         string           `json:"name" gorm:"column:name"`
	PassingGrade *float64         `json:"passing_grade,omitempty"`
	Status       *enum.QuizStatus `json:"status,omitempty"`
}

type GetStudentStatisticsDao struct {
	Student     StudentDao                    `json:"student"`
	Performance QuizPerformanceDao            `json:"performance"`
	Competency  []GetStudentQuizCompetencyDao `json:"competency"`
	Quizzes     []GetStudentQuizzesDao        `json:"quizzes"`
}
