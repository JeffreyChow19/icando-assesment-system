package dao

import "github.com/google/uuid"

type InstitutionDao struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Nis  string    `json:"nis"`
	Slug string    `json:"slug"`
}
