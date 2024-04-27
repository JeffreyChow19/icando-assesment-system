package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"icando/internal/domain/service"
	"icando/internal/worker/task"
	"icando/utils/logger"
)

type ScoreHandler struct {
	studentQuizService service.StudentQuizService
}

func NewScoreHandler(studentQuizService service.StudentQuizService) *ScoreHandler {
	return &ScoreHandler{
		studentQuizService: studentQuizService,
	}
}

func (h *ScoreHandler) HandleCalculateScoreTask() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var payload task.CalculateStudentQuizPayload
		err := json.Unmarshal(t.Payload(), &payload)

		if err != nil {
			return err
		}
		logger.Log.Info(
			fmt.Sprintf(
				"Processing calculate student quiz for id %s",
				payload.StudentQuizID.String(),
			),
		)

		err = h.studentQuizService.CalculateScore(payload.StudentQuizID)
		if err != nil {
			logger.Log.Error(err)
			return err
		}

		logger.Log.Info("Student quiz calculated successfully")
		return nil
	}
}
