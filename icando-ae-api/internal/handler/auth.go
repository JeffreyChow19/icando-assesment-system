package handler

import (
	"icando/internal/domain/service"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context, role enum.Role)
	GetTeacherProfile(c *gin.Context)
	GetStudentProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type AuthHandlerImpl struct {
	service service.AuthService
}

func NewAuthHandlerImpl(authService service.AuthService) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		service: authService,
	}
}

func (h *AuthHandlerImpl) Login(c *gin.Context, role enum.Role) {
	var loginDto dto.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authDao, err := h.service.Login(loginDto, role)

	if err != nil {
		// status error
		c.JSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, authDao)
}

func (h *AuthHandlerImpl) ChangePassword(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	userModel := user.(*dao.LearningDesignerDao)
	var changePasswordDto dto.ChangePasswordDto
	err := c.ShouldBindJSON(&changePasswordDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	httpErr := h.service.ChangePassword(userModel.ID, enum.ROLE_LEARNING_DESIGNER, changePasswordDto)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, gin.H{"error": httpErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *AuthHandlerImpl) GetTeacherProfile(c *gin.Context) {
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	claim := user.(*dao.TokenClaim)

	data, err := h.service.ProfileTeacher(claim.ID)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, data))
}

func (h *AuthHandlerImpl) GetStudentProfile(c *gin.Context) {
	studentQuiz, ok := c.Get(enum.STUDENT_QUIZ_ID_CONTEXT_KEY)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz from context"})
		return
	}
	quizData, okQuizData := studentQuiz.(*model.StudentQuiz)
	if !okQuizData {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get student quiz"})
		return
	}

	data, err := h.service.ProfileStudent(quizData.StudentID)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewBaseResponse(nil, data))
}
