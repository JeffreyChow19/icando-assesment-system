package dao

import (
	"github.com/google/uuid"
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
	Numbering  string `json:"competencyId"`
	Name       string `json:"competencyName"`
	CorrectSum int    `json:"correctCount"`
	TotalSum   int    `json:"totalCount"`
}

type GetStudentQuizzesDao struct {
	ID           uuid.UUID `json:"id"`
	QuizID       uuid.UUID `json:"quiz_id"`
	TotalScore   float32   `json:"total_score"`
	CorrectCount int       `json:"correct_count"`
	CompletedAt  time.Time `json:"completed_at"`
	Name         string    `json:"name"`
	PassingGrade float64   `json:"passing_grade"`
}

type StudentInfo struct {
	Student StudentDao `json:"student"`
	Class   ClassDao   `json:"class"`
}

type GetStudentStatisticsDao struct {
	StudentInfo StudentInfo                   `json:"student_info"`
	Performance QuizPerformanceDao            `json:"performance"`
	Competency  []GetStudentQuizCompetencyDao `json:"competency"`
	Quizzes     []GetStudentQuizzesDao        `json:"quizzes"`
}
