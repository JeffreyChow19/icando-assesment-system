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

func testRunner(t *testing.T, testFunc func(*server.Server, *lib.TestDatabase)) {
	app := fxtest.New(
		t,
		fx.Options(fx.Provide(lib.NewConfig), fx.Provide(lib.NewTestDatabase)),
		middleware.Module,
		handler.Module,
		service.Module,
		repository.Module,
		route.Module,
		server.Module,
		fx.Invoke(
			func(server *server.Server, testDb *lib.TestDatabase) {
				migrations.Up(testDb.DB)
				server.RunForTest()
				testFunc(server, testDb)
				sqlDb, _ := testDb.DB.DB()
				defer sqlDb.Close()
			},
		),
		fx.NopLogger,
	)
	defer app.RequireStop()
	app.RequireStart()
}
