package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	logOutput = "log_output"
	logLevel  = "log_level"
)
func init() {
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{getOutputLogs()},
		ErrorOutputPaths: []string{"stderr"},
	}
      var err error
	log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}
func getOutputLogs() string {
	output := strings.TrimSpace(strings.ToLower(os.Getenv(logOutput)))
	if output == "" {
		return "stdout"
	}
	return output
}
func getLevelLogs() zapcore.Level {
	switch strings.TrimSpace(strings.ToLower(os.Getenv(logLevel))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	errTags := append(tags, zap.NamedError("error", err))
	log.Error(message, errTags...)
      log.Sync()
}
