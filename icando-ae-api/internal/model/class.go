package model

import (
	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Name 	string
	Grade 	string
	InstituionID 	uuid.UUID
	TeacherID 	uuid.UUID
	Teachers 	[]Teacher
	Institution 	*Institution
}