package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Routes []Route

var Module = fx.Options(
	fx.Provide(NewRoutes),
	//fx.Provide(NewFileRoute),
	fx.Provide(NewHealthcheckRoute),
)

type Route interface {
	Setup(engine *gin.Engine)
}

func NewRoutes(
	//fileRoute *FileRoute,
	healthcheckRoute *HealthcheckRoute,
) *Routes {
	return &Routes{
		//fileRoute,
		healthcheckRoute,
	}
}

func (r Routes) Setup(engine *gin.Engine) {
	for _, route := range r {
		route.Setup(engine)
	}
}
