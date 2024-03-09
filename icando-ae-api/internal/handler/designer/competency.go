package designer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/utils/httperror"
	"net/http"
)

type CompetencyHandler interface {
	GetAllCompetencies(c *gin.Context)
	CreateCompetency(c *gin.Context)
	UpdateCompetency(c *gin.Context)
	DeleteCompetency(c *gin.Context)
}

type CompetencyHandlerImpl struct {
	competencyService service.CompetencyService
}

func NewCompetencyHandlerImpl(competencyService service.CompetencyService) *CompetencyHandlerImpl {
	return &CompetencyHandlerImpl{
		competencyService: competencyService,
	}
}

func (h *CompetencyHandlerImpl) GetAllCompetencies(c *gin.Context) {
	filter := dto.GetAllCompetenciesFilter{
		Page:  1,
		Limit: 10,
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

	competencies, meta, err := h.competencyService.GetAllCompetencies(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	competenciesResponse := &dto.GetAllCompetenciesResponse{
		Competencies: competencies,
		Meta:         *meta,
	}

	c.JSON(http.StatusOK, competenciesResponse)
}

func (h *CompetencyHandlerImpl) CreateCompetency(c *gin.Context) {
	var competency dto.CreateCompetencyDto

	if err := c.ShouldBindJSON(&competency); err != nil {
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

	competencyResponse, err := h.competencyService.CreateCompetency(competency)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, competencyResponse)
}

func (h *CompetencyHandlerImpl) UpdateCompetency(c *gin.Context) {
	var competency dto.UpdateCompetencyDto

	if err := c.ShouldBindJSON(&competency); err != nil {
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

	competencyResponse, err := h.competencyService.UpdateCompetency(competency)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, competencyResponse)
}

func (h *CompetencyHandlerImpl) DeleteCompetency(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	competencyDao, httpErr := h.competencyService.DeleteCompetency(id)
	if httpErr != nil {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, competencyDao)
}
