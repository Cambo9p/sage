package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggers = make(map[string]*zap.Logger)

func NewLogger(packageName string) (*zap.Logger, error) {
  if logger, ok := loggers[packageName]; ok {
    return logger, nil
  }

  config := zap.NewProductionConfig()
  config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
  config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

  config.OutputPaths = []string{"logs/app.log"}

  logger, err := config.Build()
  if err != nil {
    return nil, err
  }

  loggers[packageName] = logger

  return logger, nil
}
