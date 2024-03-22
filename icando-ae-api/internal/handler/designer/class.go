package designer

import (
	"errors"
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ClassHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	AssignStudents(c *gin.Context)
	UnassignStudents(c *gin.Context)
}

type ClassHandlerImpl struct {
	classService   service.ClassService
	studentService service.StudentService
}

func NewClassHandlerImpl(classService service.ClassService, studentService service.StudentService) *ClassHandlerImpl {
	return &ClassHandlerImpl{
		classService:   classService,
		studentService: studentService,
	}
}

func (h *ClassHandlerImpl) GetAll(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
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

	// check for query params withStudents, withTeacher, withInstitution
	filter := dto.GetClassFilter{}
	if withStudents, ok := c.GetQuery("withStudents"); ok {
		isWithStudents, _ := strconv.ParseBool(withStudents)

		filter.WithStudentRelation = isWithStudents
	}
	if withTeacher, ok := c.GetQuery("withTeacher"); ok {
		isWithTeacher, _ := strconv.ParseBool(withTeacher)

		filter.WithTeacherRelation = isWithTeacher
	}
	if withInstitution, ok := c.GetQuery("withInstitution"); ok {
		isWithInstitution, _ := strconv.ParseBool(withInstitution)

		filter.WithInstitutionRelation = isWithInstitution
	}

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
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstitutionID := institutionID.(uuid.UUID)

	var payload dto.CreateUpdateClassPayload

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

	student, err := h.classService.UpdateClass(parsedId, dto.ClassDto{
		Name:          payload.Name,
		Grade:         payload.Grade,
		TeacherIDs:    payload.TeacherIDs,
		InstitutionID: parsedInstitutionID,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, *student))
}

func (h *ClassHandlerImpl) Create(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstitutionID := institutionID.(uuid.UUID)

	var payload dto.CreateUpdateClassPayload

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

	student, err := h.classService.CreateClass(dto.ClassDto{
		Name:          payload.Name,
		Grade:         payload.Grade,
		TeacherIDs:    payload.TeacherIDs,
		InstitutionID: parsedInstitutionID,
	})

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

func (h *ClassHandlerImpl) AssignStudents(c *gin.Context) {
	classId := c.Param("id")
	parsedId, err := uuid.Parse(classId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}

	var studentData dto.AssignStudentsRequest

	if err := c.ShouldBindJSON(&studentData); err != nil {
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

	updatedStudents, errr := h.studentService.BatchUpdateStudentClassId(dto.UpdateStudentClassIdDto{ClassID: &parsedId, StudentIDs: studentData.StudentIDs})

	if errr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": errr})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, updatedStudents))
}

func (h *ClassHandlerImpl) UnassignStudents(c *gin.Context) {

	var studentData dto.AssignStudentsRequest

	if err := c.ShouldBindJSON(&studentData); err != nil {
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

	updatedStudents, errr := h.studentService.BatchUpdateStudentClassId(dto.UpdateStudentClassIdDto{ClassID: &uuid.Nil, StudentIDs: studentData.StudentIDs})

	if errr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": errr})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, updatedStudents))
}
