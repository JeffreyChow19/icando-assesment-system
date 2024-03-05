package dao

import "github.com/google/uuid"

type AuthDao struct {
	User  LearningDesignerDao `json:"user"`
	Token string  `json:"token"`
}

type TokenClaim struct {
	UserID uuid.UUID `json:"userId"`
	Role   string    `json:"role"`
}
