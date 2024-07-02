package scheduler

import (
	"context"
	"time"

	"github.com/HughBliss/golang_background_scheduler_example.git/internal/usecase"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/worker"
	"github.com/pkg/errors"
)

func NewScheduler(
	log logger.Logger,
	uc usecase.HelloUsecase,
) Scheduler {
	log, _ = log.Named(context.Background(), "scheduler")
	defer log.End()

	log.Info("Initializing scheduler")

	return &scheduler{
		log: log,
		uc:  uc,
	}
}

// Scheduler интерфейс для планировщика
type Scheduler interface {
	// GetPool возвращает пул воркеров
	GetPool() (worker.Pool, error)

	// sayHelloEveryMinute регистрирует задачу, которая будет выполняться каждую минуту
	sayHelloEveryMinute(payload interface{}) error
}

// scheduler реализация интерфейса Scheduler
type scheduler struct {
	log logger.Logger
	uc  usecase.HelloUsecase
}

// GetPool implements Scheduler.
func (s *scheduler) GetPool() (worker.Pool, error) {
	log, _ := s.log.Named(context.Background(), "get_pool")
	defer log.End()

	log.Info("starting register tasks in scheduler")

	pool := worker.NewPool("example", log)

	// Регистрация задачи
	if err := pool.AddTask("* * * * *", "say_hello_every_minute", s.sayHelloEveryMinute, nil); err != nil {
		log.Error("error while registering task in scheduler pool", logger.Error(err))
		return nil, errors.Wrap(err, "error registering task")
	}
	log.Info("tasks say_hello_every_minute registered in scheduler pool")

	log.Info("all tasks registered in scheduler pool")

	return pool, nil
}

// sayHelloEveryMinute implements Scheduler.
func (s *scheduler) sayHelloEveryMinute(payload interface{}) error {
	log, ctx := s.log.Named(context.Background(), "say_hello_every_minute")
	defer log.End()

	log.Info("saying hello every minute")

	if err := s.uc.SayHello(ctx, time.Now()); err != nil {
		log.Error("error while saying hello", logger.Error(err))
		return errors.Wrap(err, "error saying hello")
	}

	log.Info("hello was said")

	return nil

}
