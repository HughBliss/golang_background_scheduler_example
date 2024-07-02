package loggercore

import (
	"fmt"
	"math"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field struct {
	Key       string
	Type      FieldType
	Integer   int64
	String    string
	Interface interface{}
}

func (f *Field) ToZap() zap.Field {
	return zapcore.Field{
		Key:       f.Key,
		Type:      f.Type.ToZap(),
		Integer:   f.Integer,
		String:    f.String,
		Interface: f.Interface,
	}
}

// ToAttribute returnszapcore. a standard attribute KeyValue based on field type.
func (f *Field) ToAttribute() attribute.KeyValue {
	switch f.Type {

	case BoolType:
		val := false
		if f.Integer >= 1 {
			val = true
		}
		return attribute.Bool(f.Key, val)
	case Float32Type:
		return attribute.Float64(f.Key, math.Float64frombits(uint64(f.Integer)))
	case Float64Type:
		return attribute.Float64(f.Key, math.Float64frombits(uint64(f.Integer)))
	case Int64Type:
		return attribute.Int64(f.Key, f.Integer)
	case Int32Type:
		return attribute.Int64(f.Key, f.Integer)
	case Int16Type:
		return attribute.Int64(f.Key, f.Integer)
	case Int8Type:
		return attribute.Int64(f.Key, f.Integer)
	case Uint64Type:
		return attribute.Int64(f.Key, f.Integer)
	case Uint32Type:
		return attribute.Int64(f.Key, f.Integer)
	case Uint16Type:
		return attribute.Int64(f.Key, f.Integer)
	case Uint8Type:
		return attribute.Int64(f.Key, f.Integer)
	case StringType:
		return attribute.String(f.Key, f.String)
	case StringerType:
		return attribute.Stringer(f.Key, f.Interface.(fmt.Stringer))
	case DurationType:
		return attribute.String(f.Key, time.Duration(f.Integer).String())
	default:
		raw := fmt.Sprintf("%+v", f.Interface)
		return attribute.String(f.Key, string(raw))
	}

}

type Fields []Field

func (f Fields) ToZap() []zap.Field {
	zapfields := make([]zap.Field, 0)
	for _, field := range f {
		zapfields = append(zapfields, field.ToZap())
	}

	return zapfields
}

func (f Fields) ToAttribute() []attribute.KeyValue {
	otelAttributes := make([]attribute.KeyValue, 0)
	for _, field := range f {
		otelAttributes = append(otelAttributes, field.ToAttribute())
	}

	return otelAttributes
}
