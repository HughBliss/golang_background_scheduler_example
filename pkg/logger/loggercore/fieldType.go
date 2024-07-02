package loggercore

import "go.uber.org/zap/zapcore"

type FieldType uint8

const (
	// UnknownType is the default field type. Attempting to add it to an encoder will panic.
	UnknownType FieldType = iota
	// BinaryType indicates that the field carries an opaque binary blob.
	BinaryType
	// BoolType indicates that the field carries a bool.
	BoolType
	// ByteStringType indicates that the field carries UTF-8 encoded bytes.
	ByteStringType
	// DurationType indicates that the field carries a time.Duration.
	DurationType
	// Float64Type indicates that the field carries a float64.
	Float64Type
	// Float32Type indicates that the field carries a float32.
	Float32Type
	// Int64Type indicates that the field carries an int64.
	Int64Type
	// Int32Type indicates that the field carries an int32.
	Int32Type
	// Int16Type indicates that the field carries an int16.
	Int16Type
	// Int8Type indicates that the field carries an int8.
	Int8Type
	// StringType indicates that the field carries a string.
	StringType
	// TimeType indicates that the field carries a time.Time that is
	// representable by a UnixNano() stored as an int64.
	TimeType
	// TimeFullType indicates that the field carries a time.Time stored as-is.
	TimeFullType
	// Uint64Type indicates that the field carries a uint64.
	Uint64Type
	// Uint32Type indicates that the field carries a uint32.
	Uint32Type
	// Uint16Type indicates that the field carries a uint16.
	Uint16Type
	// Uint8Type indicates that the field carries a uint8.
	Uint8Type
	// UintptrType indicates that the field carries a uintptr.
	UintptrType
	// StringerType indicates that the field carries a fmt.Stringer.
	StringerType
	// ErrorType indicates that the field carries an error.
	ErrorType
	// SkipType indicates that the field is a no-op.
	SkipType
	// ReflectType indicates that the field carries an interface{}, which should
	// be serialized using reflection.
	ReflectType
)

// ToZap convert logger type to supported type of zap fields
func (f FieldType) ToZap() zapcore.FieldType {
	switch f {
	case UnknownType:
		return zapcore.UnknownType
	case BinaryType:
		return zapcore.BinaryType
	case BoolType:
		return zapcore.BoolType
	case ByteStringType:
		return zapcore.ByteStringType
	case DurationType:
		return zapcore.DurationType
	case Float64Type:
		return zapcore.Float64Type
	case Float32Type:
		return zapcore.Float32Type
	case Int64Type:
		return zapcore.Int64Type
	case Int32Type:
		return zapcore.Int32Type
	case Int16Type:
		return zapcore.Int16Type
	case Int8Type:
		return zapcore.Int8Type
	case StringType:
		return zapcore.StringType
	case TimeType:
		return zapcore.TimeType
	case TimeFullType:
		return zapcore.TimeFullType
	case Uint64Type:
		return zapcore.Uint64Type
	case Uint32Type:
		return zapcore.Uint32Type
	case Uint16Type:
		return zapcore.Uint16Type
	case Uint8Type:
		return zapcore.Uint8Type
	case StringerType:
		return zapcore.StringerType
	case ErrorType:
		return zapcore.ErrorType
	case SkipType:
		return zapcore.SkipType
	case ReflectType:
		return zapcore.ReflectType
	default:
		return zapcore.UnknownType
	}
}
