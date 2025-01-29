package worker

import (
	"context"
	"errors"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type WorkerFunc func(ctx context.Context, logger *zap.Logger) error

func RegisterWorker(lc fx.Lifecycle, logger *zap.Logger, shutdowner fx.Shutdowner, name string, fn WorkerFunc) {
	workerLogger := logger.With(zap.String("worker", name))
	var (
		workerCtx context.Context
		cancel    context.CancelFunc
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			workerLogger.Info("starting worker")

			workerCtx, cancel = context.WithCancel(ctx)

			go func() {
				defer func() {
					if r := recover(); r != nil {
						workerLogger.Error("worker panicked", zap.Any("error", r))
						shutdowner.Shutdown() // Gracefully stop everything
					}
				}()

				for {
					select {
					case <-workerCtx.Done():
						workerLogger.Info("worker shutting down gracefully")
						return
					default:
						if err := fn(workerCtx, workerLogger); err != nil {
							workerLogger.Error("worker failed", zap.Error(err))
							if errors.Is(err, context.Canceled) {
								return
							}

							shutdowner.Shutdown()
							return
						}
					}
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			workerLogger.Info("stopping worker")

			if cancel != nil {
				cancel()
			}

			stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer stopCancel()

			workerLogger.Info("waiting for worker to exit")
			select {
			case <-workerCtx.Done():
				workerLogger.Info("worker exited cleanly")
			case <-stopCtx.Done():
				if errors.Is(stopCtx.Err(), context.DeadlineExceeded) {
					workerLogger.Warn("worker took too long to stop, forcing shutdown")
				}
			}

			return nil
		},
	})
}
