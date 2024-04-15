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

type EmailHandler struct {
	emailService service.EmailService
}

func NewEmailHandler(emailService service.EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
	}
}

func (h *EmailHandler) HandleSendQuizEmailTask() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var payload task.SendQuizEmailPayload
		err := json.Unmarshal(t.Payload(), &payload)

		if err != nil {
			return err
		}
		logger.Log.Info(
			fmt.Sprintf(
				"Processing send quiz email task for Quiz: %s Student: %s",
				payload.QuizName,
				payload.StudentName,
			),
		)

		err = h.emailService.SendQuizEmail(payload)
		if err != nil {
			logger.Log.Error(err)
			return err
		}

		logger.Log.Info("Email sent successfully!")
		return nil
	}
}
