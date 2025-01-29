package temporalproxy

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TemporalProxyWorkerDependencies struct {
	fx.In
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
}
