package worker

import (
	"context"
	"encoding/json"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"
	"github.com/hibiken/asynq"
)

type pool struct {
	log      logger.Logger
	tasks    map[string]func(payload interface{}) error
	schedule map[*asynq.Task]string
	poolName string
}

func NewPool(poolName string, log logger.Logger) Pool {
	tasks := make(map[string]func(payload interface{}) error)
	schedule := make(map[*asynq.Task]string)
	return &pool{
		poolName: poolName,
		log:      log,
		tasks:    tasks,
		schedule: schedule,
	}
}

func (p pool) getName() string {
	return p.poolName
}
func (p pool) getSchedule() map[*asynq.Task]string {
	return p.schedule
}

func (p pool) handler(_ context.Context, t *asynq.Task) error {
	if _, ok := p.tasks[t.Type()]; !ok {
		p.log.Warnf("Task %s not found", t.Type())
		return nil
	}
	var arg interface{}
	err := json.Unmarshal(t.Payload(), &arg)
	if err != nil {
		p.log.Errorf("Error unmarshaling payload: %s", err)
		return err
	}
	err = p.tasks[t.Type()](arg)
	if err != nil {
		p.log.Errorf("Error executing task: %s", err)
		return err
	}
	return nil
}

func (p pool) AddTask(
	period,
	name string,
	handle func(payload interface{}) error,
	payload interface{},
) error {
	p.log.Infof("Adding task %s with period %s", p.poolName+":"+name, period)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		p.log.Errorf("Error marshaling payload: %s", err)
		return err
	}
	task := asynq.NewTask(p.poolName+":"+name, payloadJSON)
	p.tasks[task.Type()] = handle
	p.schedule[task] = period

	return nil
}
