package dto

import (
	"github.com/google/uuid"
	"time"
)

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GenerateQuizTokenDto struct {
	StudentQuizId uuid.UUID
	ExpiredAt     time.Time
}

type ChangePasswordDto struct {
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}
