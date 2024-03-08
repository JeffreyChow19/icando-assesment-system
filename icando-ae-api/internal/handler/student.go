package handler

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

type StudentHandler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
}

type StudentHandlerImpl struct {
	studentService service.StudentService
}

func NewStudentHandlerImpl(studentService service.StudentService) *StudentHandlerImpl {
	return &StudentHandlerImpl{
		studentService: studentService,
	}
}

func (h *StudentHandlerImpl) Post(c *gin.Context) {
	var payload dto.CreateStudentDto

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
	class, err := h.studentService.AddStudent(uuid.New(), payload)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}
	createdMsg := "Created"
	c.JSON(http.StatusCreated, response.NewBaseResponse(&createdMsg, *class))
}

func (h *StudentHandlerImpl) Get(c *gin.Context) {
	studentId := c.Param("id")
	parsedId, err := uuid.Parse(studentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("Invalid student ID").Error()})
		return
	}

	student, httperr := h.studentService.GetStudent(uuid.New(), parsedId)
	if httperr != nil {
		c.AbortWithStatusJSON(httperr.StatusCode, gin.H{"errors": httperr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *student))
}

func (h *StudentHandlerImpl) Patch(c *gin.Context) {
	studentId := c.Param("id")
	parsedId, err := uuid.Parse(studentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("Invalid student ID").Error()})
		return
	}

	var payload dto.UpdateStudentDto

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

	student, httperr := h.studentService.UpdateStudent(uuid.New(), parsedId, payload)
	if httperr != nil {
		c.AbortWithStatusJSON(httperr.StatusCode, gin.H{"errors": httperr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *student))
}

func (h *StudentHandlerImpl) Delete(c *gin.Context) {
	studentId := c.Param("id")
	parsedId, err := uuid.Parse(studentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("Invalid student ID").Error()})
		return
	}
	httperr := h.studentService.DeleteStudent(uuid.New(), parsedId)
	if httperr != nil {
		c.AbortWithStatusJSON(httperr.StatusCode, gin.H{"errors": httperr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.NewBaseResponse(nil, nil))
}
