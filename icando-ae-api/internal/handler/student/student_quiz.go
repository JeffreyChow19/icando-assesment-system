package student

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/service"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
)

type QuizHandler interface {
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

func (h *QuizHandlerImpl) UpdateAnswer(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	studentQuizID := c.Param("studentQuizId")
	questionID := c.Param("id")

	// Convert string params to uuid.UUID
	studentQuizUUID, errStudentQuizUUID := uuid.Parse(studentQuizID)
	if errStudentQuizUUID != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid studentQuizID"})
		return
	}

	questionUUID, errQuestionUUID := uuid.Parse(questionID)
	if errQuestionUUID != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid id"})
		return
	}

	var quiz dto.UpdateStudentAnswerDto
	if errBind := c.ShouldBindJSON(&quiz); errBind != nil {
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

	err := h.studentQuizService.UpdateStudentAnswer(claim.ID, studentQuizUUID, questionUUID, quiz)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	msg := "Updated"

	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, nil))
}
