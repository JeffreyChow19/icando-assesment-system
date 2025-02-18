package client

import (
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"icando/lib"
)

var Module = fx.Module("worker_client", fx.Options(fx.Provide(NewWorkerClient)))

type WorkerClient struct {
	client *asynq.Client
}

func (w *WorkerClient) Close() error {
	return w.client.Close()
}

func (w *WorkerClient) Enqueue(t *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return w.client.Enqueue(t, opts...)
}

func NewWorkerClient(config *lib.Config) *WorkerClient {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: config.RedisAddress})
	return &WorkerClient{
		client: client,
	}
}
