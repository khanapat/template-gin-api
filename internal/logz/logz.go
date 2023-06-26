package logz

import (
	"template-gin-api/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogConfig(cfg *config.Config) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "caller"

	config := zap.NewProductionConfig()
	var logLevel zapcore.Level
	switch cfg.Log.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)
	if cfg.Log.Env == "dev" {
		config.Encoding = "console"
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Encoding = "json"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
