package app

import (
	"temporal-proxy/pkg/logger"
	"temporal-proxy/pkg/workers"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func RunApplication() {
	app := fx.New(
		logger.Module,

		workers.Module,

		// Comment out to get internal FX logging back!
		fx.NopLogger,
	)

	app.Run()
}
