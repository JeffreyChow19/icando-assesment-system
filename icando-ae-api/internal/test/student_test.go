package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStudents(t *testing.T) {
	TestRunner(
		t, func(server *server.Server, db *gorm.DB, fixture Fixture) {
			students := []model.Student{
				{
					FirstName:     "Jeffrey",
					LastName:      "Chow",
					Email:         "jchow@test.com",
					InstitutionID: fixture.Instutition.ID,
					Nisn:          "1",
				},
				{
					FirstName:     "Alexander",
					LastName:      "Jason",
					Email:         "ajason@test.com",
					InstitutionID: fixture.Instutition.ID,
					Nisn:          "2",
				},
			}
			require.NoError(t, db.Create(&students).Error)
			t.Run(
				"Best Case Get All Students", func(t *testing.T) {

					w := httptest.NewRecorder()
					req, _ := CreateTeacherRequest(
						fixture, server,
						"GET",
						"/designer/student",
						nil,
					)

					server.Engine.ServeHTTP(w, req)
					require.Equal(t, http.StatusOK, w.Code)
					var res struct {
						Data []dao.StudentDao `json:"data"`
					}

					require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
					require.NotNil(t, res.Data)
					resStudents := res.Data
					require.Equal(t, 2, len(resStudents))
					require.Equal(t, students[0].FirstName, resStudents[0].FirstName)
					require.Equal(t, students[1].FirstName, resStudents[1].FirstName)
				},
			)
			t.Run(
				"Best Case Get By ID", func(t *testing.T) {

					w := httptest.NewRecorder()
					req, _ := CreateTeacherRequest(
						fixture, server,
						"GET",
						fmt.Sprintf("/designer/student/%s", students[0].ID),
						nil,
					)

					server.Engine.ServeHTTP(w, req)
					require.Equal(t, http.StatusOK, w.Code)
					var res struct {
						Data dao.StudentDao `json:"data"`
					}

					require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
					require.NotNil(t, res.Data)
					resStudents := res.Data
					require.Equal(t, students[0].FirstName, resStudents.FirstName)
				},
			)
			t.Run(
				"Filter by name", func(t *testing.T) {

					w := httptest.NewRecorder()
					req, _ := CreateTeacherRequest(
						fixture, server,
						"GET",
						"/designer/student?name=Alexander",
						nil,
					)

					server.Engine.ServeHTTP(w, req)
					require.Equal(t, http.StatusOK, w.Code)
					var res struct {
						Data []dao.StudentDao `json:"data"`
					}

					require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
					require.NotNil(t, res.Data)
					resStudents := res.Data
					require.Equal(t, 1, len(resStudents))
					require.Equal(t, students[1].FirstName, resStudents[0].FirstName)
				},
			)
		},
	)
}

func TestCreateStudent(t *testing.T) {
	TestRunner(
		t, func(server *server.Server, db *gorm.DB, fixture Fixture) {
			t.Run(
				"Best Case Create Student", func(t *testing.T) {
					payload := dto.CreateStudentDto{
						FirstName: "Alexander",
						LastName:  "Jason",
						Email:     "ajason@test.com",
						Nisn:      "2",
					}

					buff, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := CreateTeacherRequest(
						fixture, server,
						"POST",
						"/designer/student",
						bytes.NewBuffer(buff),
					)

					server.Engine.ServeHTTP(w, req)
					require.Equal(t, http.StatusCreated, w.Code)

					var students []model.Student
					require.NoError(
						t,
						db.Where("first_name = ? AND nisn = ?", payload.FirstName, payload.Nisn).Find(&students).Error,
					)
					require.Equal(t, 1, len(students))
				},
			)
			t.Run(
				"Fail create student due to empty email", func(t *testing.T) {
					payload := dto.CreateStudentDto{
						FirstName: "Alexander",
						LastName:  "Jason",
					}

					buff, _ := json.Marshal(&payload)

					w := httptest.NewRecorder()
					req, _ := CreateTeacherRequest(
						fixture, server,
						"POST",
						"/designer/student",
						bytes.NewBuffer(buff),
					)

					server.Engine.ServeHTTP(w, req)
					require.Equal(t, http.StatusBadRequest, w.Code)
				},
			)
		},
	)
}
