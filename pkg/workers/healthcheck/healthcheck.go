package healthcheck

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Dependencies managed by Fx
type HealthCheckWorkerDeps struct {
	fx.In
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
}

// HealthCheckWorker struct
type HealthCheckWorker struct {
	logger     *zap.Logger
	server     *http.Server
	shutdowner fx.Shutdowner
}

// NewHealthCheckWorker is provided by Fx
func NewHealthCheckWorker(deps HealthCheckWorkerDeps) *HealthCheckWorker {
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

var Module = fx.Options(
	fx.Provide(NewHealthCheckWorker),
	fx.Invoke(func(lifecycle fx.Lifecycle, worker *HealthCheckWorker) {
		lifecycle.Append(fx.Hook{
			OnStart: worker.Start,
			OnStop:  worker.Stop,
		})
	}),
)
