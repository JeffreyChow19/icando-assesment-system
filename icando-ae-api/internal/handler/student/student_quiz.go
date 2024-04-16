package student

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
)

type QuizHandler interface {
	StartQuiz(c *gin.Context)
	SubmitQuiz(c *gin.Context)
	UpdateAnswer(c *gin.Context)
}

type QuizHandlerImpl struct {
	studentQuizService service.StudentQuizService
}

func NewQuizHandlerImpl(studentQuizService service.StudentQuizService) *QuizHandlerImpl {
	return &QuizHandlerImpl{
		studentQuizService: studentQuizService,
	}
}

func (h *QuizHandlerImpl) StartQuiz(c *gin.Context) {
	value, ok := c.Get(enum.STUDENT_QUIZ_ID_CONTEXT_KEY)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz from context"})
		return
	}

	studentQuiz, okStudentQuiz := value.(*model.StudentQuiz)
	if !okStudentQuiz {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz"})
		return
	}

	resp, err := h.studentQuizService.StartQuiz(studentQuiz)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *resp))
}

func (h *QuizHandlerImpl) SubmitQuiz(c *gin.Context) {
	value, ok := c.Get(enum.STUDENT_QUIZ_ID_CONTEXT_KEY)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz from context"})
		return
	}

	studentQuiz, okStudentQuiz := value.(*model.StudentQuiz)
	if !okStudentQuiz {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz"})
		return
	}

	resp, err := h.studentQuizService.SubmitQuiz(studentQuiz)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *resp))
}

func (h *QuizHandlerImpl) UpdateAnswer(c *gin.Context) {
	questionID := c.Param("id")

	// Convert string params to uuid.UUID
	questionUUID, errQuestionUUID := uuid.Parse(questionID)
	if errQuestionUUID != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid id"})
		return
	}

	var req dto.UpdateStudentAnswerDto
	if errBind := c.ShouldBindJSON(&req); errBind != nil {
		var ve validator.ValidationErrors
		if errors.As(errBind, &ve) {
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

	value, ok := c.Get(enum.STUDENT_QUIZ_ID_CONTEXT_KEY)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz from context"})
		return
	}

	studentQuiz, okStudentQuiz := value.(*model.StudentQuiz)
	if !okStudentQuiz {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz"})
		return
	}

	err := h.studentQuizService.UpdateStudentAnswer(studentQuiz, questionUUID, req)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	msg := "Updated"

	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, nil))
}
