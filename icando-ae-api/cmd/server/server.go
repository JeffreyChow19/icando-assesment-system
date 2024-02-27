package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/handler"
	"icando/internal/middleware"
	"icando/internal/migrations"
	"icando/internal/route"
	"icando/lib"
	"icando/server"
	"os"
)

func startApp(server *server.Server, database *lib.Database) {
	migrations.Up(database.DB)
	server.Run()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		lib.Module,
		middleware.Module,
		handler.Module,
		service.Module,
		repository.Module,
		route.Module,
		server.Module,
		fx.Invoke(startApp),
		//fx.NopLogger, turn fx logging off
	)
	app.Run()
}
