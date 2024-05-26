package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/migrations"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/internal/route"
	"icando/internal/worker/client"
	"icando/lib"
	"icando/server"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Fixture struct {
	LearningDesigner model.Teacher
	Teacher          model.Teacher
	Instutition      model.Institution
}

func TestRunner(t *testing.T, testFunc func(*server.Server, *gorm.DB, Fixture)) {
	app := fxtest.New(
		t,
		fx.Options(fx.Provide(lib.NewConfig), fx.Provide(lib.NewDatabase)),
		fx.Invoke(
			func(db *lib.Database) {
				db.DB = db.DB.Begin()
			},
		),
		middleware.Module,
		handler.Module,
		service.Module,
		repository.Module,
		route.Module,
		server.Module,
		client.Module,
		fx.Invoke(
			func(server *server.Server, testDb *lib.Database) {
				fixture := SetupFixture(testDb.DB)
				migrations.Up(testDb.DB)
				server.RunForTest()
				testFunc(server, testDb.DB, fixture)
				sqlDb, _ := testDb.DB.DB()
				defer testDb.DB.Rollback()
				defer sqlDb.Close()
			},
		),
		fx.NopLogger,
	)
	defer app.RequireStop()
	app.RequireStart()
}

func SetupFixture(db *gorm.DB) Fixture {
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	instutition := model.Institution{
		Name: "Test123",
		Slug: "test123",
		Nis:  "1234",
	}
	db.Create(&instutition)
	teacher := model.Teacher{
		Email:         "learning_designer@test.com",
		Password:      string(pw),
		Role:          enum.TEACHER_ROLE_LEARNING_DESIGNER,
		InstitutionID: instutition.ID,
	}
	db.Create(&teacher)

	return Fixture{
		LearningDesigner: teacher,
		Teacher:          teacher,
		Instutition:      instutition,
	}
}

func CreateTeacherRequest(f Fixture, server *server.Server, method string, url string, payload io.Reader) (
	*http.Request, error,
) {
	loginPayload := dto.LoginDto{
		Email:    f.Teacher.Email,
		Password: "password",
	}
	body, _ := json.Marshal(loginPayload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/auth/designer/login",
		bytes.NewBuffer(body),
	)
	server.Engine.ServeHTTP(w, req)

	var res dao.AuthDao
	json.Unmarshal(w.Body.Bytes(), &res)

	req, err := http.NewRequest(method, url, payload)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", res.Token))
	return req, err
}
