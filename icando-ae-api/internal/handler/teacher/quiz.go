package teacher

import (
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type QuizHandler interface {
	GetAllQuizDetail(c *gin.Context)
	GetStudentQuiz(c *gin.Context)
}

type QuizHandlerImpl struct {
	quizDetailService    service.QuizService
	studentQuizService   service.StudentQuizService
	competencyRepository repository.CompetencyRepository
}

func NewQuizHandlerImpl(
	quizDetailService service.QuizService,
	studentQuizService service.StudentQuizService,
	competencyRepository repository.CompetencyRepository,
) *QuizHandlerImpl {
	return &QuizHandlerImpl{
		quizDetailService:    quizDetailService,
		studentQuizService:   studentQuizService,
		competencyRepository: competencyRepository,
	}
}

func (h *QuizHandlerImpl) GetAllQuizDetail(c *gin.Context) {
	institutionID, _ := c.Get(enum.INSTITUTION_ID_CONTEXT_KEY)
	parsedInstutionID := institutionID.(uuid.UUID).String()

	filter := dto.GetAllQuizzesFilter{
		InstitutionID: &parsedInstutionID,
		Page:          1,
		Limit:         10,
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

	// todo?: filter if quiz is teached by teacher
	quizzes, meta, err := h.quizDetailService.GetAllQuizzes(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponseWithMeta(&createdMsg, quizzes, meta))
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
