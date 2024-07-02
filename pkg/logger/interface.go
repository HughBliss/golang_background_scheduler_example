package logger

import (
	"context"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger/loggercore"
	"go.opentelemetry.io/otel/trace"
)

type Logger interface {
	Debug(msg string, fields ...loggercore.Field)
	Info(msg string, fields ...loggercore.Field)
	Warn(msg string, fields ...loggercore.Field)
	Error(msg string, fields ...loggercore.Field)
	Fatal(msg string, fields ...loggercore.Field)
	Panic(msg string, fields ...loggercore.Field)

	// Formating message
	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)
	Panicf(template string, args ...any)
	End()
	Ctx() context.Context
	Named(ctx context.Context, name string, opt ...trace.SpanStartOption) (Logger, context.Context)
}
