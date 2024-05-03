package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/httperror"
	"icando/utils/response"
	"net/http"
)

type QuizHandler interface {
	GetAllQuizDetail(c *gin.Context)
	GetQuizHistory(c *gin.Context)
	GetStudentQuiz(c *gin.Context)
	GetStudentQuizzes(c *gin.Context)
	GetQuiz(c *gin.Context)
}

type QuizHandlerImpl struct {
	studentQuizService   service.StudentQuizService
	quizService          service.QuizService
	competencyRepository repository.CompetencyRepository
}

func NewQuizHandlerImpl(
	studentQuizService service.StudentQuizService,
	competencyRepository repository.CompetencyRepository,
	quizService service.QuizService,
) *QuizHandlerImpl {
	return &QuizHandlerImpl{
		studentQuizService:   studentQuizService,
		competencyRepository: competencyRepository,
		quizService:          quizService,
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

	quizzes, meta, err := h.quizService.GetAllQuizzes(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponseWithMeta(&createdMsg, quizzes, meta))
}

func (h *QuizHandlerImpl) GetQuizHistory(c *gin.Context) {
	quizID := c.Param("id")
	parsedQuizID, err := uuid.Parse(quizID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid class ID").Error()})
		return
	}
	filter := dto.GetQuizVersionFilter{
		ID:    parsedQuizID,
		Page:  1,
		Limit: 10,
	}

	quizHistory, meta, httpErr := h.quizService.GetQuizHistory(filter)

	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, gin.H{"errors": httpErr.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponseWithMeta(&createdMsg, quizHistory, meta))

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

	studentQuizCompetency, err := h.competencyRepository.GetStudentCompetency(
		dto.GetStudentCompetencyFilter{
			StudentQuizID: &quiz.ID,
		},
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(
		http.StatusOK, response.NewBaseResponse(
			nil, map[string]any{
				"quiz":       *quiz,
				"competency": studentQuizCompetency,
			},
		),
	)
}

func (h *QuizHandlerImpl) GetQuiz(c *gin.Context) {
	quizId := c.Param("quizId")
	parsedId, err := uuid.Parse(quizId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors.New("invalid quiz ID").Error()})
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

	quiz, httperr := h.quizService.GetQuiz(
		dto.GetQuizFilter{
			ID: parsedId, WithCreator: true, WithUpdater: true,
			WithQuestions: true, WithClasses: true,
		},
	)

	if httperr != nil {
		c.AbortWithStatusJSON(httperr.StatusCode, gin.H{"errors": httperr.Err.Error()})
		return
	}
	flattenedClassIds := []uuid.UUID{}
	for _, class := range quiz.Classes {
		flattenedClassIds = append(flattenedClassIds, class.ID)
	}
	IsTeachingClass, err := teacher.IsTeachingClasses(flattenedClassIds)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	if !IsTeachingClass {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": errors.New("You do not teach this class")})
		return
	}

	studentQuizCompetency, err := h.competencyRepository.GetStudentCompetency(
		dto.GetStudentCompetencyFilter{
			StudentQuizID: &quiz.ID,
		},
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(
		http.StatusOK, response.NewBaseResponse(
			nil, map[string]any{
				"quiz":       *quiz,
				"competency": studentQuizCompetency,
			},
		),
	)
}

func (h *QuizHandlerImpl) GetStudentQuizzes(c *gin.Context) {
	var filter dto.GetStudentQuizzesFilter

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

	if filter.QuizID != nil {
		_, errQuizUUID := uuid.Parse(*filter.QuizID)
		if errQuizUUID != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid quiz id"})
			return
		}
	}
	teacherId := teacher.ID.String()
	filter.TeacherID = &teacherId
	studentStatistics, meta, err := h.studentQuizService.GetStudentQuizzes(filter)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"errors": err.Err.Error()})
		return
	}

	createdMsg := "ok"
	c.JSON(http.StatusOK, response.NewBaseResponseWithMeta(&createdMsg, studentStatistics, meta))
}
