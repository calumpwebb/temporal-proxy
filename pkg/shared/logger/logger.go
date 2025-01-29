package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	// TODO: (calum) prod logger?
	// TODO: (calum) tag all logs etc
	// TODO: (calum) metrics\
	print("WE GOT THE LOGGER!")
	return zap.NewDevelopment()
}
