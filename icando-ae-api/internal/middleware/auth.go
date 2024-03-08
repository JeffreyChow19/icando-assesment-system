package middleware

import (
	"fmt"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthMiddleware struct {
	studentRepository          repository.StudentRepository
	teacherRepository          repository.TeacherRepository
	learningDesignerRepository repository.LearningDesignerRepository
	config                     *lib.Config
}

func NewAuthMiddleware(
	config *lib.Config,
	studentRepository repository.StudentRepository,
	teacherRepository repository.TeacherRepository,
	learningDesignerRepository repository.LearningDesignerRepository,
) *AuthMiddleware {
	return &AuthMiddleware{
		studentRepository:          studentRepository,
		teacherRepository:          teacherRepository,
		learningDesignerRepository: learningDesignerRepository,
		config:                     config,
	}
}

func (m *AuthMiddleware) Handler(role model.Role) gin.HandlerFunc {
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

			if authorized.Role == model.ROLE_TEACHER {
				teacher, err := m.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &authorized.ID})

				c.Set("InstitutionID", teacher.InstitutionID)

				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errors.New("Teacher not found")})
					return
				}
			} else if authorized.Role == model.ROLE_STUDENT {
				_, err := m.studentRepository.GetOne(dto.GetStudentFilter{ID: &authorized.ID})

				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errors.New("Student not found")})
					return
				}
			} else if authorized.Role == model.ROLE_LEARNING_DESIGNER {
				_, err := m.learningDesignerRepository.FindLearningDesigner(dto.GetLearningDesignerFilter{ID: &authorized.ID})

				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errors.New("Learning designer not found")})
					return
				}
			}

			if authorized.Role == role || role == model.ROLE_ALL {
				c.Set("user", authorized)
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (m *AuthMiddleware) authorize(tokenString string) (*dao.TokenClaim, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.config.JwtSecret), nil
	})

	if !token.Valid || err != nil {
		return nil, errors.New("token is invalid")
	}

	exp := claims["exp"].(float64)

	if exp <= float64(time.Now().Unix()) {
		return nil, errors.New("token expired")
	}

	id := fmt.Sprint(claims["id"])

	parsedUuid, err := uuid.Parse(id)

	role := fmt.Sprint(claims["role"])

	if err != nil {
		return nil, errors.New("Cannot parse uuid")
	}

	var parsedRole model.Role

	if role == model.ROLE_LEARNING_DESIGNER {
		parsedRole = model.ROLE_LEARNING_DESIGNER
	} else if role == model.ROLE_STUDENT {
		parsedRole = model.ROLE_STUDENT
	} else if role == model.ROLE_TEACHER {
		parsedRole = model.ROLE_TEACHER
	}

	return &dao.TokenClaim{
		ID:   parsedUuid,
		Role: parsedRole,
	}, nil
}
