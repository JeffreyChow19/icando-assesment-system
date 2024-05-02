package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/response"
	"net/http"
)

type QuizHandler interface {
	GetStudentQuiz(c *gin.Context)
}

type QuizHandlerImpl struct {
	studentQuizService   service.StudentQuizService
	competencyRepository repository.CompetencyRepository
}

func NewQuizHandlerImpl(
	studentQuizService service.StudentQuizService,
	competencyRepository repository.CompetencyRepository,
) *QuizHandlerImpl {
	return &QuizHandlerImpl{
		studentQuizService:   studentQuizService,
		competencyRepository: competencyRepository,
	}
}

func (h *QuizHandlerImpl) GetStudentQuiz(c *gin.Context) {
	studentQuizId := c.Param("studentquizid")
	parsedId, err := uuid.Parse(studentQuizId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid student quiz ID").Error()})
		return
	}

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

	quiz, httperr := h.studentQuizService.GetQuizDetailByID(parsedId)

	if httperr != nil {
		c.AbortWithStatusJSON(httperr.StatusCode, gin.H{"errors": httperr.Err.Error()})
		return
	}

	IsTeachingClass, err := teacher.IsTeachingClass(*quiz.Student.ClassID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	if !IsTeachingClass {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": errors.New("You do not teach this class")})
		return
	}

	studentQuizCompetency, err := h.competencyRepository.GetStudentCompetency(dto.GetStudentCompetencyFilter{
		StudentQuizID: &quiz.ID,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, map[string]any{
		"quiz":       *quiz,
		"competency": studentQuizCompetency,
	}))
}
