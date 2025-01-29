package healthcheck

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HealthCheckWorkerDependencies struct {
	fx.In
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
}
