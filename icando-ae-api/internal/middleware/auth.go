package middleware

import (
	"fmt"
	"icando/internal/domain/service"
	"icando/lib"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthMiddleware struct{
	service service.LearningDesignerService
	config *lib.Config
}

func NewAuthMiddleware(service service.LearningDesignerService, config *lib.Config) *AuthMiddleware{
	return &AuthMiddleware{
		service: service,
		config: config,
	}
}

func (m *AuthMiddleware) Handler(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 && t[0] == "Bearer" {
			authToken := t[1]
			authorized, err := m.authorize(authToken)
			if authorized == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			user, httpError := m.service.FindUserById(*authorized)
			if httpError != nil {
				c.AbortWithStatusJSON(httpError.StatusCode, gin.H{"error": httpError.Err.Error()})
				return
			}
			
			if user.Role == role  {
				c.Set("user", user)
				c.Next()
				return
			}			
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (m *AuthMiddleware) authorize(tokenString string) (*uuid.UUID, error){
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
		return []byte(m.config.JwtSecret), nil
	},)

	if(!token.Valid || err != nil){
		return nil, errors.New("token is invalid")
	}
	exp := claims["exp"].(float64)
	if exp <= float64(time.Now().Unix()) {
		return nil, errors.New("token expired")
	}

	id := fmt.Sprint(claims["id"])

	parsedUuid, _ := uuid.Parse(id)
	return &parsedUuid, nil
}