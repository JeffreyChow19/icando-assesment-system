package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/worker"
	"icando/internal/worker/handler"
	"icando/lib"
	"icando/utils/logger"
	"os"
)

// consumer server
func startServer(server *worker.WorkerServer) {
	logger.Log.Info("Starting worker server")
	server.Run()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		lib.Module,
		service.Module,
		repository.Module,
		handler.Module,
		worker.Module,
		fx.Invoke(startServer),
		fx.NopLogger,
	)
	app.Run()
}
