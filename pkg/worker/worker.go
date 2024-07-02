package worker

import (
	"context"
	"sync"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/config"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"

	"github.com/hibiken/asynq"
)

type worker struct {
	log       logger.Logger
	server    *asynq.Server
	mux       *asynq.ServeMux
	scheduler *asynq.Scheduler
	pools     map[string]Pool
}

func NewWorker(
	log logger.Logger,
	cfg *config.WorkerConfig,
) (Worker, error) {
	log, _ = log.Named(context.Background(), "worker")
	defer log.End()
	log.Info("Initializing worker")
	workerLogger := NewWorkerLogger(log)
	redisConnOpt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		DB:       cfg.Redis.Db,
		Password: cfg.Redis.Password,
	}
	server := asynq.NewServer(redisConnOpt, asynq.Config{Concurrency: 10,
		Queues: map[string]int{
			"critical": cfg.CriticalQueuePriority,
			"default":  cfg.DefaultQueuePriority,
			"low":      cfg.LowQueuePriority,
		},
		Logger: workerLogger},
	)
	scheduler := asynq.NewScheduler(redisConnOpt, &asynq.SchedulerOpts{Logger: workerLogger})
	mux := asynq.NewServeMux()
	pools := make(map[string]Pool)
	return &worker{
		log:       log,
		pools:     pools,
		server:    server,
		mux:       mux,
		scheduler: scheduler,
	}, nil
}

func (p worker) bindTasks(schedule map[*asynq.Task]string) error {
	for task, period := range schedule {
		ID, err := p.scheduler.Register(period, task)
		if err != nil {
			p.log.Errorf("Error registering task: %s", err)
			return err
		}
		p.log.Infof("Task %s registered with ID %s", task.Type(), ID)
	}
	return nil
}

func (p worker) AddPool(pool Pool) {
	p.mux.HandleFunc(pool.getName(), pool.handler)
	p.pools[pool.getName()] = pool
}

func (p worker) Serve() error {

	for _, pool := range p.pools {
		err := p.bindTasks(pool.getSchedule())
		if err != nil {
			return err
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	go func() {
		defer wg.Done()
		if err = p.scheduler.Run(); err != nil {
			p.log.Errorf("Error running scheduler: %s", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err = p.server.Run(p.mux); err != nil {
			p.log.Fatalf("Start worker, err: %s", err)
		}
	}()
	wg.Wait()
	return err
}
