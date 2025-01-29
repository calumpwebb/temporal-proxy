package workers

import (
	"temporal-proxy/pkg/workers/healthcheck"
	"temporal-proxy/pkg/workers/temporalproxy"

	"go.uber.org/fx"
)

var Module = fx.Options(
	healthcheck.Module,
	temporalproxy.Module,
)
