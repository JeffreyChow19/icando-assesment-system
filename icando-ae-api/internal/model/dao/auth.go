package dao

import (
	"github.com/google/uuid"
	"icando/internal/model/enum"
)

type AuthDao struct {
	User  TokenClaim `json:"user"`
	Token string     `json:"token"`
}

type TokenClaim struct {
	ID   uuid.UUID `json:"id"`
	Role enum.Role `json:"role"`
	Exp  int64     `json:"exp"`
}
