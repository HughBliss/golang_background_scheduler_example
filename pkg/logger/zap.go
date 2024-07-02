package logger

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger/loggercore"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TODO: Add work with context like here
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go
type ZapLogger struct {
	logger      *zap.Logger
	cfg         *zap.Config
	innerCtx    context.Context
	name        string
	serivceName string
	endCh       chan struct{}
}

var _ Logger = (*ZapLogger)(nil)

const rootSreviceName = "root"

// TODO: Add method new logger from context
func New(cfg *Config) (Logger, error) {
	endCh := make(chan struct{})

	zapConfig, err := cfg.Log.ZapConfig()
	if err != nil {
		return nil, err
	}
	zaplog, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	zaplog = zaplog.Named(cfg.ServiceName)
	tp, err := NewJaegerTraceProvider(cfg)
	if err != nil {
		zaplog.Fatal("exporter create error", zap.Error(err))
	}

	otel.SetTracerProvider(tp)
	rootCtx, span := otel.Tracer(cfg.ServiceName).Start(context.Background(), rootSreviceName)
	span.SetAttributes(
		attribute.String("message", "create root logger"),
		attribute.String("name", rootSreviceName),
	)

	go shutdown(rootCtx, tp, endCh)

	return &ZapLogger{
		logger:      zaplog,
		cfg:         zapConfig,
		name:        rootSreviceName,
		serivceName: cfg.ServiceName,
		innerCtx:    rootCtx,
		endCh:       endCh,
	}, nil

}

func (l *ZapLogger) span(ctx context.Context, msg string, fields ...loggercore.Field) {
	span := trace.SpanFromContext(ctx)

	otelFields := make([]attribute.KeyValue, 0)
	otelFields = append(otelFields, attribute.String("level", l.cfg.Level.String()))
	if msg != "" {
		otelFields = append(otelFields, attribute.String("message", msg))
	}

	if len(fields) > 0 {
		otelFields = append(otelFields, loggercore.Fields(fields).ToAttribute()...)
	}
	span.AddEvent(l.name, trace.WithAttributes(otelFields...))
}

func (l *ZapLogger) Debug(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Debug(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Info(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Info(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Warn(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Warn(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Error(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Error(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) DPanic(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.DPanic(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Panic(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Panic(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Fatal(msg string, fields ...loggercore.Field) {
	l.span(l.innerCtx, msg, fields...)
	l.logger.Fatal(msg, loggercore.Fields(fields).ToZap()...)
}

func (l *ZapLogger) Debugf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Debugf(msg, args)
}

func (l *ZapLogger) Infof(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Infof(msg, args...)
}

func (l *ZapLogger) Warnf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Warnf(msg, args...)
}

func (l *ZapLogger) Errorf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Errorf(msg, args...)
}

func (l *ZapLogger) DPanicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().DPanicf(msg, args...)
}

func (l *ZapLogger) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Panicf(msg, args...)
}

func (l *ZapLogger) Fatalf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.span(l.innerCtx, msg)
	l.logger.Sugar().Fatalf(msg, args...)
}

// If call on root logger, shutdown TraceProvider
func (l *ZapLogger) End() {
	span := trace.SpanFromContext(l.innerCtx)
	span.End()

	if l.name == rootSreviceName {
		l.endCh <- struct{}{}
	}
}

func (l *ZapLogger) Ctx() context.Context {
	return l.innerCtx
}

func (l *ZapLogger) Named(ctx context.Context, name string, opt ...trace.SpanStartOption) (Logger, context.Context) {
	named := new(ZapLogger)
	named.logger = l.logger.Named(name)
	named.cfg = l.cfg
	named.serivceName = l.serivceName
	named.name = strings.Join([]string{l.name, name}, ".")
	newCtx, span := otel.Tracer(l.serivceName).Start(ctx, name, opt...)

	named.innerCtx = newCtx
	span.SetAttributes(
		attribute.String("message", "create logger"),
		attribute.String("name", named.name),
	)
	return named, newCtx
}

func shutdown(ctx context.Context, tp *tracesdk.TracerProvider, endCh chan struct{}) {
	<-endCh
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		panic(err)
	}
}
