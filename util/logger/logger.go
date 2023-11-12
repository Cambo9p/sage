package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggers = make(map[string]*zap.Logger)

// this logger doesnt handle errors well, we jsut ignore them as making the parent
// handle them is just plain annoying and increases checks
func NewLogger(packageName string) *zap.Logger {
	if logger, ok := loggers[packageName]; ok {
		return logger
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	config.OutputPaths = []string{"logs/app.log"}

	logger, err := config.Build()
	if err != nil {
		return nil
	}

	loggers[packageName] = logger

	return logger
}
