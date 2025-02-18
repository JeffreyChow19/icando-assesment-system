package designer

import (
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TeacherHandler interface {
	GetAll(c *gin.Context)
}

type TeacherHandlerImpl struct {
	teacherService service.TeacherService
}

func NewTeacherHandlerImpl(teacherService service.TeacherService) *TeacherHandlerImpl {
	return &TeacherHandlerImpl{
		teacherService: teacherService,
	}
}

func (h *TeacherHandlerImpl) GetAll(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstitutionID := institutionID.(uuid.UUID)

	filter := dto.GetTeacherFilter{
		InstitutionID: &parsedInstitutionID,
	}

	teachers, err := h.teacherService.GetAllTeachers(filter)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, teachers))
}
