package dao

import (
	"github.com/google/uuid"
)

type AuthDao struct {
	User  TokenClaim `json:"user"`
	Token string     `json:"token"`
}

type TokenClaim struct {
	ID  uuid.UUID `json:"id"`
	Exp int64     `json:"exp"`
}
