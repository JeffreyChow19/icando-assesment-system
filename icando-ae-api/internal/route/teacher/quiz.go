package teacher

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/teacher"
)

type QuizRoute struct {
	quizHandler teacher.QuizHandler
}

func (r QuizRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/quiz")
	group.GET("/:quizid/students/:studentquizid", r.quizHandler.GetStudentQuiz)
}

func NewQuizRoute(
	handler teacher.QuizHandler,
) *QuizRoute {
	return &QuizRoute{
		quizHandler: handler,
	}
}
