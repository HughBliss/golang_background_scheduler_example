package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level string `json:"level" yaml:"level"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development" yaml:"development"`
	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string `json:"encoding" yaml:"encoding"`
}

type TelemetryConfig struct {
	Host string `yaml:"host" env:"TELEMETRY_HOST" env-default:"0.0.0.0"`
	Port string `yaml:"port" env:"TELEMETRY_PORT" env-default:"6831"`
}

type Config struct {
	ServiceName string          `yaml:"service_name" env:"SERVICE_NAME" env-default:"payments_microservice"`
	ServiceVer  string          `yaml:"service_ver" env:"SERVICE_VER" env-default:"0.0.1"`
	Telemetry   TelemetryConfig `yaml:"telemetry"`
	Log         LogConfig       `yaml:"logger"`
}

func (c LogConfig) ZapConfig() (*zap.Config, error) {
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      c.Development,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         c.Encoding,
		InitialFields:    c.InitialFields,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "name",
			StacktraceKey: "stack",
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
		},
	}, nil
}
