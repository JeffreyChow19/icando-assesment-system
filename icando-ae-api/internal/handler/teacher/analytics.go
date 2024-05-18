package teacher

import (
	"errors"
	"github.com/google/uuid"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dao"
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
	GetDashboardOverview(c *gin.Context)
	GetAllStudents(c *gin.Context)
	GetAllClasses(c *gin.Context)
}

type AnalyticsHandlerImpl struct {
	analyticsService service.AnalyticsService
	studentService   service.StudentService
	classService     service.ClassService
}

func NewAnalyticsHandlerImpl(analyticsService service.AnalyticsService, studentService service.StudentService, classService service.ClassService) *AnalyticsHandlerImpl {
	return &AnalyticsHandlerImpl{
		analyticsService: analyticsService,
		studentService:   studentService,
		classService:     classService,
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

	filter := dto.GetLatestSubmissionsFilter{
		TeacherID: &teacherID,
		Page:      1,
		Limit:     5,
	}

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

	latestSubmissions, err := h.analyticsService.GetLatestSubmissions(filter)
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

func (h *AnalyticsHandlerImpl) GetDashboardOverview(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	data, err := h.analyticsService.GetDashboardOverview(claim.ID)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, data))
}

func (h *AnalyticsHandlerImpl) GetAllStudents(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstutionID := institutionID.(uuid.UUID).String()

	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)
	teacherId := claim.ID.String()

	filter := dto.GetAllStudentsFilter{
		InstitutionID: &parsedInstutionID,
		Page:          1,
		Limit:         10,
		TeacherID:     &teacherId,
	}

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

	students, meta, err := h.studentService.GetAllStudents(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponseWithMeta(&createdMsg, students, meta))
}

func (h *AnalyticsHandlerImpl) GetAllClasses(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstitutionID := institutionID.(uuid.UUID)
	sortBy := "name"

	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	filter := dto.GetAllClassFilter{
		TeacherID:     &claim.ID,
		InstitutionID: &parsedInstitutionID,
		SortBy:        &sortBy,
	}

	class, err := h.classService.GetAllClass(filter)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, class))
}
