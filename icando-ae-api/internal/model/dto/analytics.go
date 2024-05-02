package dto

type GetQuizPerformanceFilter struct {
	StudentID *string `form:"studentId"`
	TeacherID *string `form:"teacherId"`
	QuizID    *string `form:"quizId"`
}
