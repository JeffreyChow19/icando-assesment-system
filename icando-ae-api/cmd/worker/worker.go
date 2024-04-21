package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"icando/internal/domain/repository"
	"icando/internal/domain/service"
	"icando/internal/worker/client"
	"icando/internal/worker/handler"
	"icando/internal/worker/server"
	"icando/lib"
	"icando/utils/logger"
	"os"
)

// consumer server
func startServer(server *server.WorkerServer) {
	logger.Log.Info("Starting worker server")
	server.Run()
}

func NewWorkerClient() *client.WorkerClient {
	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		lib.Module,
		fx.Options(fx.Provide(NewWorkerClient)),
		service.Module,
		repository.Module,
		handler.Module,
		server.Module,
		fx.Invoke(startServer),
		// fx.NopLogger,
	)
	app.Run()
}
