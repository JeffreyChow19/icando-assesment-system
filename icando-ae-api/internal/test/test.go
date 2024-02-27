package test

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/migrations"
	"icando/internal/route"
	"icando/lib"
	"icando/server"
	"testing"
)

func testRunner(t *testing.T, testFunc func(*server.Server)) {
	app := fxtest.New(
		t, lib.Module,
		middleware.Module,
		handler.Module,
		service.Module,
		repository.Module,
		route.Module,
		server.Module,
		fx.Invoke(
			func(server *server.Server, testDb *lib.Database) {
				migrations.Up(testDb.DB)
				server.RunForTest()
				testFunc(server)
				sqlDb, _ := testDb.DB.DB()
				defer sqlDb.Close()
			},
		),
		fx.NopLogger,
	)
	defer app.RequireStop()
	app.RequireStart()
}
