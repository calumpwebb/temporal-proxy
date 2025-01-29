package healthcheck

import (
	"context"
	"errors"
	"net/http"
	"temporal-proxy/pkg/worker"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// HealthCheckWorker runs a simple HTTP health check server
func HealthCheckWorker(ctx context.Context, logger *zap.Logger) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

		logger.Info("/healthz called")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to capture server start errors
	serverErr := make(chan error, 1)

	// Start HTTP server in a goroutine
	go func() {
		logger.Info("server started on " + server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err // Send error to main loop
		}
	}()

	// Periodic log to confirm worker is alive
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("server stopping")
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				logger.Error("forced shutdown", zap.Error(err))
			} else {
				logger.Info("server stopped cleanly")
			}
			return nil

		case err := <-serverErr:
			// If server fails to start, log error and return it to shut down fx
			logger.Error("server startup failed", zap.Error(err))
			return err

		case <-ticker.C:
			logger.Info("server running")
		}
	}
}

// Register worker in fx lifecycle
var invoke = fx.Invoke(func(lc fx.Lifecycle, logger *zap.Logger, shutdowner fx.Shutdowner) {
	worker.RegisterWorker(lc, logger, shutdowner, "HealthCheckWorker", HealthCheckWorker)
	// worker.RegisterWorker(lc, logger, shutdowner, "HealthCheckWorker", HealthCheckWorker)
})

var Module = fx.Options(
	invoke,
)
