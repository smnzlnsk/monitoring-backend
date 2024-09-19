package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func NewLogger() error {
	l, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	Logger = l
	return nil
}
