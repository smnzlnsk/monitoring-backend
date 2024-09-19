package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func NewLogger() error {
	config := zap.NewDevelopmentConfig()

	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	l, err := config.Build()
	if err != nil {
		return err
	}
	Logger = l
	return nil
}
