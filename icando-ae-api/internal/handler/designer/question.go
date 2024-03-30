package designer

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
)

type QuestionHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type QuestionHandlerImpl struct {
	questionService service.QuestionService
}

func NewQuestionHandlerImpl(questionService service.QuestionService) *QuestionHandlerImpl {
	return &QuestionHandlerImpl{
		questionService: questionService,
	}
}

func (h *QuestionHandlerImpl) Create(c *gin.Context) {
	quizID := c.Param("quizId")

	parsedQuizID, err := uuid.Parse(quizID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid quizId"})
		return
	}

	var question dto.QuestionDto
	if err := c.ShouldBindJSON(&question); err != nil {
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

	questionResponse, errResponse := h.questionService.CreateQuestion(parsedQuizID, question)

	if errResponse != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": errResponse})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *questionResponse))
}

func (h *QuestionHandlerImpl) Update(c *gin.Context) {
	quizID := c.Param("quizId")
	questionID := c.Param("id")

	// Convert string params to uuid.UUID
	quizUUID, err := uuid.Parse(quizID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid quizId"})
		return
	}

	questionUUID, err := uuid.Parse(questionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid id"})
		return
	}

	filter := dto.GetQuestionFilter{
		ID:     questionUUID,
		QuizID: quizUUID,
	}

	var question dto.QuestionDto
	if err := c.ShouldBindJSON(&question); err != nil {
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

	questionResponse, errResponse := h.questionService.UpdateQuestion(filter, question)

	if errResponse != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": errResponse})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *questionResponse))
}

func (h *QuestionHandlerImpl) Delete(c *gin.Context) {
	quizID := c.Param("quizId")
	questionID := c.Param("id")

	quizUUID, err := uuid.Parse(quizID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid quizId"})
		return
	}

	questionUUID, errQuestionUUID := uuid.Parse(questionID)
	if errQuestionUUID != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid id"})
		return
	}

	filter := dto.GetQuestionFilter{
		ID:     questionUUID,
		QuizID: quizUUID,
	}

	errDeleteQuestion := h.questionService.DeleteQuestion(filter)

	if errDeleteQuestion != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": errDeleteQuestion})
		return
	}

	msg := "Deleted"

	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, nil))
}
