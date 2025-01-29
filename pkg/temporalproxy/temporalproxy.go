package temporalproxy

import (
	"context"
	"temporal-proxy/pkg/worker"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func TemporalProxyWorker(ctx context.Context, logger *zap.Logger) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("I'm stopping")
			return nil
		case <-ticker.C:
			logger.Info("I'm running...")
			panic("test")
		}
	}
}

var invoke = fx.Invoke(func(lc fx.Lifecycle, logger *zap.Logger, shutdowner fx.Shutdowner) {
	worker.RegisterWorker(lc, logger, shutdowner, "TemporalProxyWorker", TemporalProxyWorker)
})

var Module = fx.Options(
	invoke,
)
