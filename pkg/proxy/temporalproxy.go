package proxy

import (
	"context"
	"temporal-proxy/pkg/internal/service"
	internalservice "temporal-proxy/pkg/internal/service"
	"time"
)

type TemporalProxy struct {
	serviceRunner service.ServiceRunner
}

func NewTemporalProxy(id string) *TemporalProxy {
	temporalProxyService := newTemporalProxyService()

	defaultServiceRunnerOptions := internalservice.NewDefaultServiceRunnerOptions(
		service.WithShutdownTimeout(30 * time.Second),
	)

	serviceRunner := internalservice.NewDefaultServiceRunner(temporalProxyService, defaultServiceRunnerOptions)

	return &TemporalProxy{
		serviceRunner: serviceRunner,
	}
}

func (tp TemporalProxy) Run(ctx context.Context) {
	tp.serviceRunner.Run(ctx)
}
