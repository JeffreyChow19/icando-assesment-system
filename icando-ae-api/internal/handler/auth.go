package handler

import (
	"icando/internal/domain/service"
	"icando/internal/model/dto"
	"icando/internal/model/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler interface {
	Login(c *gin.Context)
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

func (h *AuthHandlerImpl) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authDao, err := h.service.Login(loginDto)
	if err != nil {
		// status error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authDao)
}

func (h *AuthHandlerImpl) ChangePassword(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*dao.LearningDesignerDao)
	var changePasswordDto dto.ChangePasswordDto
	err := c.ShouldBindJSON(&changePasswordDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	httpErr := h.service.ChangePassword(userModel.ID, changePasswordDto)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, gin.H{"error": httpErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
