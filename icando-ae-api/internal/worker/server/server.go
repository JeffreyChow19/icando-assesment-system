package server

import (
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"icando/internal/worker/handler"
	"icando/internal/worker/task"
	"icando/lib"
	"log"
)

type WorkerServer struct {
	srv    *asynq.Server
	router *asynq.ServeMux
	db     *gorm.DB
}

func NewServer(config *lib.Config, db *lib.Database, emailHandler *handler.EmailHandler) *WorkerServer {
	var server WorkerServer
	redisConnOpt := asynq.RedisClientOpt{Addr: config.RedisAddress}
	server.srv = asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	server.router = asynq.NewServeMux()
	server.router.HandleFunc(task.TypeSendQuizEmailTask, emailHandler.HandleSendQuizEmailTask())
	server.db = db.DB
	return &server
}

func (s *WorkerServer) Run() {
	if err := s.srv.Run(s.router); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

var Module = fx.Module("worker_server", fx.Options(fx.Provide(NewServer)))
