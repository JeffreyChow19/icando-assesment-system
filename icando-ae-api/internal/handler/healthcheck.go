package handler

import (
	"github.com/gin-gonic/gin"
	"icando/internal/model/enum"
	"icando/utils/response"
	"net/http"
)

type HealthcheckHandler interface {
	Healthcheck(c *gin.Context)
	HealthcheckProtected(c *gin.Context)
}

type HealthcheckHandlerImpl struct {
}

func (h *HealthcheckHandlerImpl) Healthcheck(c *gin.Context) {
	msg := "🚀Service up and running🚀"
	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, nil))
}

func (h *HealthcheckHandlerImpl) HealthcheckProtected(c *gin.Context) {
	msg := "🚀Protected🚀"
	user, _ := c.Get(enum.USER_CONTEXT_KEY)
	c.JSON(http.StatusOK, response.NewBaseResponse(&msg, user))
}

func NewHealthcheckHandlerImpl() *HealthcheckHandlerImpl {
	return &HealthcheckHandlerImpl{}
}
