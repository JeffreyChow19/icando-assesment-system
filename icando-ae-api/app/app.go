package app

import (
	"go.uber.org/fx"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/migrations"
	"icando/internal/route"
	"icando/lib"
	"icando/server"
)

func RunServer(server *server.Server, database *lib.Database) {
	migrations.Up(database.DB)
	server.Run()
}

func StartApp() {
	app := fx.New(
		lib.Module,
		middleware.Module,
		handler.Module,
		service.Module,
		repository.Module,
		route.Module,
		server.Module,
		fx.Invoke(RunServer),
		//fx.NopLogger, turn fx logging off
	)
	app.Run()
}
