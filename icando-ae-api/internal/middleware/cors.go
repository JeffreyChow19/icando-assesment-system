package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"icando/lib"
	"time"
)

func CorsMiddleware(config *lib.Config) gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins: []string{config.ClientHost},
			AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
			AllowHeaders: []string{
				"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "sec-ch-ua",
			},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}
