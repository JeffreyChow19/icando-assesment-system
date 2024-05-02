package middleware

import (
	"fmt"
	"icando/internal/domain/repository"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
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
	studentRepository     repository.StudentRepository
	teacherRepository     repository.TeacherRepository
	studentQuizRepository repository.StudentQuizRepository
	config                *lib.Config
}

func NewAuthMiddleware(
	config *lib.Config,
	studentRepository repository.StudentRepository,
	teacherRepository repository.TeacherRepository,
	studentQuizRepository repository.StudentQuizRepository,
) *AuthMiddleware {
	return &AuthMiddleware{
		studentRepository:     studentRepository,
		teacherRepository:     teacherRepository,
		studentQuizRepository: studentQuizRepository,
		config:                config,
	}
}

func (m *AuthMiddleware) Handler(role enum.Role) gin.HandlerFunc {
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

			if role == enum.ROLE_TEACHER || role == enum.ROLE_LEARNING_DESIGNER {
				// :<
				trueVal := true
				teacher, err := m.teacherRepository.GetTeacher(dto.GetTeacherFilter{ID: &authorized.ID, WithClasses: &trueVal})

				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errors.New("Teacher not found")})
					return
				}

				if role == enum.ROLE_LEARNING_DESIGNER && teacher.Role != enum.ROLE_LEARNING_DESIGNER {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errors.New("Forbidden")})
					return
				}

				c.Set(enum.INSTITUTION_ID_CONTEXT_KEY, teacher.InstitutionID)
				c.Set(enum.USER_CONTEXT_KEY, authorized)
				c.Set(enum.TEACHER_CONTEXT_KEY, teacher)
			} else if role == enum.ROLE_STUDENT {
				idString := authorized.ID.String()

				idUUID, err := uuid.Parse(idString)

				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errors.New("Cannot parse uuid")})
					return
				}

				studentQuiz, err := m.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{ID: idUUID})

				if err != nil {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errors.New("Student quiz not found")})
					return
				}

				c.Set(enum.STUDENT_QUIZ_ID_CONTEXT_KEY, studentQuiz)
				c.Set(enum.USER_CONTEXT_KEY, dao.TokenClaim{
					ID:  studentQuiz.StudentID,
					Exp: authorized.Exp,
				})
			}

			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (m *AuthMiddleware) authorize(tokenString string) (*dao.TokenClaim, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.config.JwtSecret), nil
		},
	)

	if !token.Valid || err != nil {
		return nil, errors.New("token is invalid")
	}

	exp := claims["exp"].(float64)

	if exp <= float64(time.Now().Unix()) {
		return nil, errors.New("token expired")
	}

	id := fmt.Sprint(claims["id"])

	parsedUuid, err := uuid.Parse(id)

	if err != nil {
		return nil, errors.New("Cannot parse uuid")
	}

	return &dao.TokenClaim{
		ID: parsedUuid,
	}, nil
}
