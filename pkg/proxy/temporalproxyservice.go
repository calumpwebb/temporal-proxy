package proxy

import (
	"context"
	"fmt"
	"sync/atomic"
	"temporal-proxy/pkg/internal/service"
	"time"
)

type temporalProxyService struct {
	running atomic.Bool
}

func newTemporalProxyService() service.Service {
	return &temporalProxyService{}
}

func (s *temporalProxyService) Run(ctx context.Context) error {
	s.running.Store(true)

	for s.running.Load() {
		select {
		case <-ctx.Done(): // Listen for cancellation
			fmt.Println("TemporalProxyService: received stop signal")
			return nil
		default:
			fmt.Println("TemporalProxyService: running")
			time.Sleep(time.Second)
		}
	}

	return nil
}

func (s *temporalProxyService) Stop() {
	fmt.Println("TemporalProxyService: stopping")
	s.running.Store(false)

	for {
		time.Sleep(time.Second)
		fmt.Println("cont")
	}
}
