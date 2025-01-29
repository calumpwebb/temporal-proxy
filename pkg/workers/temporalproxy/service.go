package temporalproxy

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TemporalProxyWorker struct {
	logger     *zap.Logger
	shutdowner fx.Shutdowner
}

func NewTemporalProxyWorker(dependencies TemporalProxyWorkerDependencies) *TemporalProxyWorker {
	return &TemporalProxyWorker{
		logger:     dependencies.Logger,
		shutdowner: dependencies.Shutdowner,
	}
}

func (h *TemporalProxyWorker) Start(ctx context.Context) error {
	h.logger.Info("starting temporal proxy server")

	// TODO: (calum) actually implemen this here!!

	// Run server in a goroutine so Start() doesn't block
	go func() {
		// TODO: (calum) i think this ctx gets cancelled after 15 seconds... new context required? or smoething else, why not an issue for healthcheck?
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				h.logger.Info("temporal proxy server closing")
				// TODO: (calum) shutdowner? what about if this is an error? feels like this class is stillclosign for some reason
				h.logger.Info(ctx.Err().Error())
				return
			case <-ticker.C:
				h.logger.Info("temporal proxy server is running")
			}
		}
	}()

	h.logger.Info("temporal proxy server started")

	return nil
}

func (h *TemporalProxyWorker) Stop(ctx context.Context) error {
	h.logger.Info("shutting down temporal proxy server")
	return nil
}
