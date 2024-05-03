package teacher

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/teacher"
)

type StudentQuizRoute struct {
	quizHandler teacher.QuizHandler
}

func (r StudentQuizRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/student-quiz")
	group.GET("/:studentquizid", r.quizHandler.GetStudentQuiz)
	group.GET("", r.quizHandler.GetStudentQuizzes)
}

func NewStudentQuizRoute(
	handler teacher.QuizHandler,
) *StudentQuizRoute {
	return &StudentQuizRoute{
		quizHandler: handler,
	}
}
