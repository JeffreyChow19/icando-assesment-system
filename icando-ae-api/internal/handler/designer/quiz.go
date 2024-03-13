package designer

import (
	"github.com/gin-gonic/gin"
	"icando/internal/domain/service"
	"icando/internal/model/dao"
	"icando/internal/model/enum"
	"icando/utils/response"
	"net/http"
)

type QuizHandler interface {
	Create(c *gin.Context)
}

type QuizHandlerImpl struct {
	quizService service.QuizService
}

func NewQuizHandlerImpl(quizService service.QuizService) *QuizHandlerImpl {
	return &QuizHandlerImpl{
		quizService: quizService,
	}
}

func (h *QuizHandlerImpl) Create(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	quiz, err := h.quizService.CreateQuiz(claim.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *quiz))
}
