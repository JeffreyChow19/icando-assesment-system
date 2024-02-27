package route

import (
	"github.com/gin-gonic/gin"
	"icando/internal/handler"
)

type FileRoute struct {
	fileHandler handler.FileHandler
}

func (r FileRoute) Setup(engine *gin.Engine) {
	group := engine.Group("/file")
	group.POST("/image", r.fileHandler.RequestImageUpload)
	//group.POST("/invalidate", r.authMiddleware.Handler(model.ROLE_ADMIN), r.fileHandler.InvalidateCache)
}

func NewFileRoute(handler handler.FileHandler) *FileRoute {
	return &FileRoute{
		fileHandler: handler,
		//authMiddleware: authMiddleware,
	}
}
