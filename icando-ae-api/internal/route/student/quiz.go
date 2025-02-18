package student

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler/student"
	"icando/internal/middleware"
)

type QuizRoute struct {
	quizHandler    student.QuizHandler
	authMiddleware middleware.AuthMiddleware
}

func (r QuizRoute) Setup(group *gin.RouterGroup) {
	group = group.Group("/quiz")
	group.GET("/overview", r.quizHandler.GetQuiz)
	group.GET("/detail", r.quizHandler.GetQuizDetail)
	group.POST("/question/:id", r.quizHandler.UpdateAnswer)
	group.PATCH("/start", r.quizHandler.StartQuiz)
	group.PATCH("/submit", r.quizHandler.SubmitQuiz)
	group.GET("/review", r.quizHandler.GetQuizReview)
}

func NewQuizRoute(
	handler student.QuizHandler,
	authMiddleware *middleware.AuthMiddleware,
) *QuizRoute {
	return &QuizRoute{
		quizHandler:    handler,
		authMiddleware: *authMiddleware,
	}
}
