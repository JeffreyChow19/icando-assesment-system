package handler

import (
	"github.com/gin-gonic/gin"
)

type FileHandler interface {
	RequestImageUpload(c *gin.Context)
	RequestNostalgiaUpload(c *gin.Context)
	RequestNostalgiaBlackUpload(c *gin.Context)
	InvalidateCache(c *gin.Context)
}
