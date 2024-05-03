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
	ClassName   string    `json:"className"`
	Grade       string    `json:"grade"`
	QuizName    string    `json:"quizName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CompletedAt time.Time `json:"completedAt"`
}

type GetStudentQuizCompetencyDao struct {
	Numbering  string `json:"competencyId"`
	Name       string `json:"competencyName"`
	CorrectSum int    `json:"correctCount"`
	TotalSum   int    `json:"totalCount"`
}

type GetStudentQuizzesDao struct {
	ID           uuid.UUID `json:"id"`
	QuizID       uuid.UUID `json:"quizId"`
	TotalScore   float32   `json:"totalScore"`
	CorrectCount int       `json:"correctCount"`
	CompletedAt  time.Time `json:"completedAt"`
	Name         string    `json:"name"`
	PassingGrade float64   `json:"passingGrade"`
}

type StudentInfo struct {
	Student StudentDao `json:"student"`
	Class   ClassDao   `json:"class"`
}

type GetStudentStatisticsDao struct {
	StudentInfo StudentInfo                   `json:"studentInfo"`
	Performance QuizPerformanceDao            `json:"performance"`
	Competency  []GetStudentQuizCompetencyDao `json:"competency"`
	Quizzes     []GetStudentQuizzesDao        `json:"quizzes"`
}
