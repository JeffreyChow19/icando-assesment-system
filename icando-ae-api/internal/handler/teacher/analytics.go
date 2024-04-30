package teacher

import (
	"errors"
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnalyticsHandler interface {
	GetQuizPerformance(c *gin.Context)
}

type AnalyticsHandlerImpl struct {
	analyticsService service.AnalyticsService
}

func NewAnalyticsHandlerImpl(analyticsService service.AnalyticsService) *AnalyticsHandlerImpl {
	return &AnalyticsHandlerImpl{
		analyticsService: analyticsService,
	}
}

func (h *AnalyticsHandlerImpl) GetQuizPerformance(c *gin.Context) {
	// todo: institution ID di filter ???

	filter := dto.GetQuizPerformanceFilter{}

	if err := c.ShouldBindQuery(&filter); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]httperror.FieldError, len(ve))
			for i, fe := range ve {
				out[i] = httperror.FieldError{Field: fe.Field(), Message: httperror.MsgForTag(fe.Tag()), Tag: fe.Tag()}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": "Invalid query"})
		return
	}

	quizPerformance, err := h.analyticsService.GetQuizPerformance(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponse(&createdMsg, quizPerformance))
}
