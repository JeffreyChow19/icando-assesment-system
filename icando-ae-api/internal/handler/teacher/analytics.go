package teacher

import (
	"errors"
	"github.com/google/uuid"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnalyticsHandler interface {
	GetQuizPerformance(c *gin.Context)
	GetLatestSubmissions(c *gin.Context)
	GetStudentStatistics(c *gin.Context)
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

func (h *AnalyticsHandlerImpl) GetLatestSubmissions(c *gin.Context) {
	value, ok := c.Get(enum.TEACHER_CONTEXT_KEY)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get teacher from context"})
		return
	}

	teacher, okTeacher := value.(*model.Teacher)

	if !okTeacher {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get teacher data"})
		return
	}

	teacherID := teacher.ID.String()

	latestSubmissions, err := h.analyticsService.GetLatestSubmissions(dto.GetLatestSubmissionsFilter{
		TeacherID: &teacherID,
	})
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponse(&createdMsg, latestSubmissions))
}

func (h *AnalyticsHandlerImpl) GetStudentStatistics(c *gin.Context) {
	studentID := c.Param("id")

	// Convert string params to uuid.UUID
	studentUUID, errStudentUUID := uuid.Parse(studentID)
	if errStudentUUID != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid id"})
		return
	}

	studentStatistics, err := h.analyticsService.GetStudentStatistics(studentUUID)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponse(&createdMsg, studentStatistics))
}
