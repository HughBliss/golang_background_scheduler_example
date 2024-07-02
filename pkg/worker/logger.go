package worker

import "github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"

type WorkerLogger interface {
	// Debug logs a message at Debug level.
	Debug(args ...interface{})

	// Info logs a message at Info level.
	Info(args ...interface{})

	// Warn logs a message at Warning level.
	Warn(args ...interface{})

	// Error logs a message at Error level.
	Error(args ...interface{})

	// Fatal logs a message at Fatal level
	// and process will exit with status set to 1.
	Fatal(args ...interface{})
}

type workerLogger struct {
	log logger.Logger
}

func (w workerLogger) Debug(args ...interface{}) {
	w.log.Debugf("%v", args)
}

func (w workerLogger) Info(args ...interface{}) {
	w.log.Infof("%v", args)
}

func (w workerLogger) Warn(args ...interface{}) {
	w.log.Warnf("%v", args)
}

func (w workerLogger) Error(args ...interface{}) {
	w.log.Errorf("%v", args)
}

func (w workerLogger) Fatal(args ...interface{}) {
	w.log.Fatalf("%v", args)
}

func NewWorkerLogger(log logger.Logger) WorkerLogger {
	return &workerLogger{log: log}
}
