package app

import (
	"temporal-proxy/pkg/healthcheck"
	"temporal-proxy/pkg/logger"
	"temporal-proxy/pkg/temporalproxy"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func RunApplication() {
	app := fx.New(
		logger.Module,
		temporalproxy.Module,
		healthcheck.Module,

		// Comment out to get internal FX logging back!
		fx.NopLogger,
	)

	app.Run()
}
