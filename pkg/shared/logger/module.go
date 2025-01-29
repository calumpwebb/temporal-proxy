package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("logger",
	fx.Provide(NewLogger),

	// Use the same logger for internal Fx logging
	fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		println("still got here")
		println(logger)
		return &fxevent.ZapLogger{Logger: logger}
	}),
)
