package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"icando/internal/model/base"
	"time"
)

const (
	TypeSendQuizEmailTask = "email:quiz"
)

type SendQuizEmailPayload struct {
	QuizName     string
	QuizSubjects base.StringArray
	QuizDuration int
	QuizEndAt    time.Time
	QuizStartAt  time.Time
	QuizUrl      string
	TeacherName  string
	TeacherEmail string
	StudentName  string
	StudentEmail string
}

func NewSendQuizEmailTask(quizDetails SendQuizEmailPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(quizDetails)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendQuizEmailTask, payload, asynq.MaxRetry(3)), nil
}
