package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin_LearningDesigner(t *testing.T) {
	TestRunner(
		t, func(server *server.Server, db *gorm.DB, fixture Fixture) {
			t.Run(
				"Best Case learning designer logn", func(t *testing.T) {
					payload := dto.LoginDto{
						Email:    fixture.LearningDesigner.Email,
						Password: "password",
					}
					body, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := http.NewRequest(
						"POST",
						"/auth/designer/login",
						bytes.NewBuffer(body),
					)

					server.Engine.ServeHTTP(w, req)

					require.Equal(t, http.StatusOK, w.Code)

					var res dao.AuthDao
					json.Unmarshal(w.Body.Bytes(), &res)
					require.NotEmpty(t, res.User)
					require.NotEmpty(t, res.Token)
				},
			)
			t.Run(
				"Wrong password case", func(t *testing.T) {
					payload := dto.LoginDto{
						Email:    fixture.LearningDesigner.Email,
						Password: "password123",
					}
					body, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := http.NewRequest(
						"POST",
						"/auth/designer/login",
						bytes.NewBuffer(body),
					)

					server.Engine.ServeHTTP(w, req)

					require.Equal(t, http.StatusUnauthorized, w.Code)
				},
			)
		},
	)
}

func TestLogin_Teacher(t *testing.T) {
	TestRunner(
		t, func(server *server.Server, db *gorm.DB, fixture Fixture) {
			t.Run(
				"Best Case learning designer logn", func(t *testing.T) {
					payload := dto.LoginDto{
						Email:    fixture.Teacher.Email,
						Password: "password",
					}
					body, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := http.NewRequest(
						"POST",
						"/auth/teacher/login",
						bytes.NewBuffer(body),
					)

					server.Engine.ServeHTTP(w, req)

					require.Equal(t, http.StatusOK, w.Code)

					var res dao.AuthDao
					json.Unmarshal(w.Body.Bytes(), &res)
					require.NotEmpty(t, res.User)
					require.NotEmpty(t, res.Token)
				},
			)
			t.Run(
				"Wrong password case", func(t *testing.T) {
					payload := dto.LoginDto{
						Email:    fixture.Teacher.Email,
						Password: "password123",
					}
					body, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := http.NewRequest(
						"POST",
						"/auth/teacher/login",
						bytes.NewBuffer(body),
					)

					server.Engine.ServeHTTP(w, req)

					require.Equal(t, http.StatusUnauthorized, w.Code)
				},
			)
		},
	)
}
