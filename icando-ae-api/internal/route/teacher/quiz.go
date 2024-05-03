package teacher

import (
	"icando/internal/handler/teacher"

	"github.com/gin-gonic/gin"
)

type QuizRoute struct {
	quizHandler teacher.QuizHandler
}

func (r QuizRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/quiz")
	group.GET("", r.quizHandler.GetAllQuizDetail)
	group.GET("/:quizid/students/:studentquizid", r.quizHandler.GetStudentQuiz)
}

func NewQuizRoute(
	handler teacher.QuizHandler,
) *QuizRoute {
	return &QuizRoute{
		quizHandler: handler,
	}
}
