package task

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	TypeCalculateStudentQuizTask = "quiz:score"
)

type CalculateStudentQuizPayload struct {
	StudentQuizID uuid.UUID
}

func NewCalcualteStudentQuizTask(data CalculateStudentQuizPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCalculateStudentQuizTask, payload, asynq.MaxRetry(3)), nil
}
