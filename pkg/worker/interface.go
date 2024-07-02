package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type (
	Pool interface {
		AddTask(
			period,
			name string,
			handle func(payload interface{}) error,
			payload interface{},
		) error
		handler(ctx context.Context, t *asynq.Task) error
		getName() string
		getSchedule() map[*asynq.Task]string
	}
	Worker interface {
		Serve() error
		AddPool(pool Pool)
		bindTasks(schedule map[*asynq.Task]string) error
	}
)
