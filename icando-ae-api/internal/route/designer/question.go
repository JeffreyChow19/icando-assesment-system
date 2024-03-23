package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/designer"
)

type QuestionRoute struct {
	questionHandler designer.QuestionHandler
}

func (r QuestionRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/quiz")
	group.POST(":quizId/question", r.questionHandler.Create)
	group.PATCH(":quizId/question/:id", r.questionHandler.Update)
	group.DELETE(":quizId/question/:id", r.questionHandler.Delete)
}

func NewQuestionRoute(
	handler designer.QuestionHandler,
) *QuestionRoute {
	return &QuestionRoute{
		questionHandler: handler,
	}
}
