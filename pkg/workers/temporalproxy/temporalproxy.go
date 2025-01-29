package temporalproxy

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TemporalProxyWorkerDeps struct {
	fx.In
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
}

type TemporalProxyWorker struct {
	logger     *zap.Logger
	shutdowner fx.Shutdowner
}

func NewTemporalProxyWorker(deps TemporalProxyWorkerDeps) *TemporalProxyWorker {
	logger := deps.Logger

	return &TemporalProxyWorker{
		logger:     logger,
		shutdowner: deps.Shutdowner,
	}
}

func (h *TemporalProxyWorker) Start(ctx context.Context) error {
	h.logger.Info("starting temporal proxy server")

	// Run server in a goroutine so Start() doesn't block
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				h.logger.Info("temporal proxy server closing")
				return
			case <-ticker.C:
				h.logger.Info("temporal proxy server is running")
			}
		}
	}()

	return nil
}

func (h *TemporalProxyWorker) Stop(ctx context.Context) error {
	h.logger.Info("shutting down temporal proxy server")
	return nil
}

var Module = fx.Options(
	fx.Provide(NewTemporalProxyWorker),
	fx.Invoke(func(lifecycle fx.Lifecycle, worker *TemporalProxyWorker) {
		lifecycle.Append(fx.Hook{
			OnStart: worker.Start,
			OnStop:  worker.Stop,
		})
	}),
)
