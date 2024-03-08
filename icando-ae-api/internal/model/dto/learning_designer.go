package dto

import "github.com/google/uuid"

type PutUserInfoDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type GetLearningDesignerFilter struct {
	ID    *uuid.UUID
	Email *string
}
