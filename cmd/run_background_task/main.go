package main

import (
	"log"

	"github.com/HughBliss/golang_background_scheduler_example.git/internal/config"
	"github.com/HughBliss/golang_background_scheduler_example.git/internal/scheduler"
	"github.com/HughBliss/golang_background_scheduler_example.git/internal/usecase"
	pkgConfig "github.com/HughBliss/golang_background_scheduler_example.git/pkg/config"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/worker"
)

func main() {

	cfg := config.Get()

	rootLogger, err := logger.New(&logger.Config{
		ServiceName: cfg.ServiceName,
		ServiceVer:  cfg.ServiceVer,
		Telemetry:   cfg.Telemetry,
		Log:         cfg.Log,
	})
	if err != nil {
		log.Fatal(err)
	}

	rootLogger.Info("Starting application")

	w, err := worker.NewWorker(rootLogger, &pkgConfig.WorkerConfig{
		Redis:                 &cfg.Redis,
		CriticalQueuePriority: 6,
		DefaultQueuePriority:  3,
		LowQueuePriority:      1,
	})
	if err != nil {
		rootLogger.Fatal("error while creating worker", logger.Error(err))
		return
	}

	uc := usecase.NewHelloUsecase(rootLogger)

	s := scheduler.NewScheduler(rootLogger, uc)

	pool, err := s.GetPool()
	if err != nil {
		rootLogger.Fatal("error while getting pool", logger.Error(err))
		return
	}

	w.AddPool(pool)

	if err := w.Serve(); err != nil {
		rootLogger.Fatal("error while serving worker", logger.Error(err))
		return
	}
}
