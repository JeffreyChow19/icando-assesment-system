package student

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/student"
)

type QuizRoute struct {
	quizHandler student.QuizHandler
}

func (r QuizRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/quiz")
	group.POST(":studentQuizId/question/:id", r.quizHandler.UpdateAnswer)
}

func NewQuizRoute(
	handler student.QuizHandler,
) *QuizRoute {
	return &QuizRoute{
		quizHandler: handler,
	}
}
