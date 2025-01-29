package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type DefaultServiceRunner struct {
	service Service
	options DefaultServiceRunnerOptions
}

func NewDefaultServiceRunner(service Service, options DefaultServiceRunnerOptions) ServiceRunner {
	return &DefaultServiceRunner{
		service: service,
		options: options,
	}
}

func (dsr *DefaultServiceRunner) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Handle OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := dsr.service.Run(ctx); err != nil {
			fmt.Printf("Service encountered an error: %v\n", err)
		}
	}()

	// Wait for termination signal
	<-stop
	fmt.Println("Received termination signal. Shutting down...")

	// Request shutdown
	dsr.service.Stop()

	// Wait for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), dsr.options.ShutdownTimeout)
	defer shutdownCancel()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Service shut down gracefully")
	case <-shutdownCtx.Done():
		fmt.Println("Shutdown timeout exceeded. Force quitting.")
		os.Exit(1)
	}
}
