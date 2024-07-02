package logger

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// newResource returns a resource describing this application.
func newResource(cfg *Config) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(cfg.ServiceName),
			semconv.ServiceVersion(cfg.ServiceVer),
			attribute.String("environment", "debug"),
		),
	)
	return r
}

// NewJaegerTraceProvider returns new trace provider with jaeger exporter
func NewJaegerTraceProvider(cfg *Config) (*otelsdk.TracerProvider, error) {
	telemetryCfg := cfg.Telemetry
	exp, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(telemetryCfg.Host),
		jaeger.WithAgentPort(telemetryCfg.Port),
	))

	if err != nil {
		return nil, err
	}

	return otelsdk.NewTracerProvider(
		otelsdk.WithBatcher(exp),
		otelsdk.WithResource(newResource(cfg)),
	), nil
}
