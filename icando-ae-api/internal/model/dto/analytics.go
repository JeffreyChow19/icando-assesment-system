package dto

import "icando/internal/model/enum"

type GetQuizPerformanceFilter struct {
	StudentID *string `form:"studentId"`
	TeacherID *string `form:"teacherId"`
	QuizID    *string `form:"quizId"`
}

type GetLatestSubmissionsFilter struct {
	TeacherID *string `form:"teacher_id"`
}

type GetStudentQuizzesFilter struct {
	QuizID      *string          `form:"quizId"`
	ClassID     *string          `form:"classId"`
	StudentName *string          `form:"studentName"`
	QuizStatus  *enum.QuizStatus `form:"quizStatus"`
	Page        int              `form:"page"`
	Limit       int              `form:"limit"`
	TeacherID   *string
}
