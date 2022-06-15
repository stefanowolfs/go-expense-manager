package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = "" // removes print stacktrace on Error logs
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Info receives a message and logs it at the InfoLevel.
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Debug receives a message and logs it at the DebugLevel.
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Error receives a message and logs it at the ErrorLevel.
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
