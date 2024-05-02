package dao

import "time"

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
