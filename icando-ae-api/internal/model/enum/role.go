package enum

type Role string

const (
	ROLE_LEARNING_DESIGNER = "LEARNING_DESIGNER"
	ROLE_TEACHER           = "TEACHER"
	ROLE_STUDENT           = "STUDENT"
	ROLE_ALL               = "ALL"
)

type TeacherRole string

const (
	TEACHER_ROLE_REGULAR           = "REGULAR"
	TEACHER_ROLE_LEARNING_DESIGNER = "LEARNING_DESIGNER"
)
