package dto

import "github.com/google/uuid"

type GetManyInstitutionFilter struct {
}

type GetOneInstitutionFilter struct {
	ID   *uuid.UUID
	Nis  *string
	Slug *string
}
