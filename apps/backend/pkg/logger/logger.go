package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the global logger instance
var Log *zap.Logger

func Init(env string) {
	var cfg zap.Config

	if env == "production" {
		// Production: JSON format, info level, no dev extras
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// Development: human-readable, debug level, colour
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var err error
	Log, err = cfg.Build()
	if err != nil {
		panic("failed to initialise logger: " + err.Error())
	}
}

// Sync flushes buffered log entries — call on shutdown
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
