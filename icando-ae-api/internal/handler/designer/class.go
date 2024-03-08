package designer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
)

type ClassHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type ClassHandlerImpl struct {
	classService service.ClassService
}

func NewClassHandlerImpl(classService service.ClassService) *ClassHandlerImpl {
	return &ClassHandlerImpl{
		classService: classService,
	}
}

func (h *ClassHandlerImpl) GetAll(c *gin.Context) {
	institutionID, _ := c.Get("InstitutionID")
	parsedInstitutionID := institutionID.(uuid.UUID)
	sortBy := "name"

	filter := dto.GetAllClassFilter{
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

func (h *ClassHandlerImpl) Get(c *gin.Context) {
	classId := c.Param("id")
	parsedId, err := uuid.Parse(classId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}

	// change param here to include other relation
	filter := dto.GetClassFitler{}

	class, err := h.classService.GetClass(parsedId, filter)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *class))
}

func (h *ClassHandlerImpl) Update(c *gin.Context) {
	classId := c.Param("id")
	parsedId, err := uuid.Parse(classId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}

	var payload dto.ClassDto

	if err := c.ShouldBindJSON(&payload); err != nil {
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

	student, err := h.classService.UpdateClass(parsedId, payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *student))
}

func (h *ClassHandlerImpl) Create(c *gin.Context) {
	var payload dto.ClassDto

	if err := c.ShouldBindJSON(&payload); err != nil {
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

	student, err := h.classService.CreateClass(payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *student))
}

func (h *ClassHandlerImpl) Delete(c *gin.Context) {
	classId := c.Param("id")
	parsedId, err := uuid.Parse(classId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}

	err = h.classService.DeleteClass(parsedId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	msg := "Deleted"

	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, nil))
}
