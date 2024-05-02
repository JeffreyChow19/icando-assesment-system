package dto

type GetQuizPerformanceFilter struct {
	StudentID *string `form:"studentId"`
	TeacherID *string `form:"teacherId"`
	QuizID    *string `form:"quizId"`
}

type GetLatestSubmissionsFilter struct {
	TeacherID *string `form:"teacher_id"`
}
