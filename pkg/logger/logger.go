package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

var Module = fx.Options(
	fx.Provide(NewLogger),

	// Use the same logger for internal Fx logging
	fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger}
	}),
)
