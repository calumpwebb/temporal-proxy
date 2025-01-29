package app

import (
	"temporal-proxy/pkg/shared/logger"
	"temporal-proxy/pkg/workers/healthcheck"
	"temporal-proxy/pkg/workers/temporalproxy"

	"go.uber.org/fx"
)

var modules = fx.Module("app",
	logger.Module,

	healthcheck.Module,
	temporalproxy.Module,

	// Comment out to get internal FX logging back!
	// fx.NopLogger,
)
