package healthcheck

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HealthCheckWorker struct {
	logger     *zap.Logger
	server     *http.Server
	shutdowner fx.Shutdowner
}

func NewHealthCheckWorker(deps HealthCheckWorkerDependencies) *HealthCheckWorker {
	logger := deps.Logger

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

		logger.Info("GET /healthz check endpoint called")
	})

	server := &http.Server{
		Addr:    ":8080", // TODO: (calum) configurable?
		Handler: mux,
	}

	return &HealthCheckWorker{
		logger:     logger,
		server:     server,
		shutdowner: deps.Shutdowner,
	}
}

func (h *HealthCheckWorker) Start(ctx context.Context) error {
	h.logger.Info("starting health check server on " + h.server.Addr)

	// TODO: (calum) i dont think this will shut down if ctx is cancelled... its never passed in? (maybe its okay because its shutdown in Stop)
	// Run server in a goroutine so Start() doesn't block
	go func() {
		if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.logger.Error("health check server failed", zap.Error(err))

			// shut down the rest of the app because we failed to start
			h.shutdowner.Shutdown()
		}
	}()

	return nil
}

func (h *HealthCheckWorker) Stop(ctx context.Context) error {
	h.logger.Info("shutting down health check server")

	shutdownCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return h.server.Shutdown(shutdownCtx)
}
