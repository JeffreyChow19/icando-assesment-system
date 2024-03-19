package designer

import (
	"icando/internal/domain/service"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type QuizHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
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

func (h *QuizHandlerImpl) Get(c *gin.Context) {
	quizId := c.Param("id")
	parsedId, err := uuid.Parse(quizId)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}
	
	quiz, errr := h.quizService.GetQuiz(parsedId)

	if errr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *quiz))
}

func (h *QuizHandlerImpl) Update(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	var quiz dto.UpdateQuizDto
	if err := c.ShouldBindJSON(&quiz); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]httperror.FieldError, len(ve))
			for i, fe := range ve {
				out[i] = httperror.FieldError{Field: fe.Field(), Message: httperror.MsgForTag(fe.Tag()), Tag: fe.Tag()}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": "Invalid body"})
		return
	}

	quizResponse, err := h.quizService.UpdateQuiz(claim.ID, quiz)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *quizResponse))
}
