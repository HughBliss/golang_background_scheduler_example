package logger

import (
	"context"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger/loggercore"

	"go.opentelemetry.io/otel/trace"
)

// TODO: Add work with context like here
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go
type EmptyLogger struct {
}

var _ Logger = (*EmptyLogger)(nil)

// TODO: Add method new logger from context
func Empty() Logger {

	return &EmptyLogger{}
}

func (l *EmptyLogger) Debug(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Info(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Warn(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Error(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) DPanic(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Panic(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Fatal(msg string, fields ...loggercore.Field) {
}

func (l *EmptyLogger) Debugf(template string, args ...any) {
}

func (l *EmptyLogger) Infof(template string, args ...any) {
}

func (l *EmptyLogger) Warnf(template string, args ...any) {
}

func (l *EmptyLogger) Errorf(template string, args ...any) {
}

func (l *EmptyLogger) DPanicf(template string, args ...any) {
}

func (l *EmptyLogger) Panicf(template string, args ...any) {
}

func (l *EmptyLogger) Fatalf(template string, args ...any) {
}

// If call on root logger, shutdown TraceProvider
func (l *EmptyLogger) End() {
}

func (l *EmptyLogger) Ctx() context.Context {
	return context.Background()
}

func (l *EmptyLogger) Named(ctx context.Context, name string, opt ...trace.SpanStartOption) (Logger, context.Context) {

	return l, context.Background()
}
