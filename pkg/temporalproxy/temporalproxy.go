package temporalproxy

import (
	"errors"
	"sync/atomic"
	"temporal-proxy/pkg/service"

	"go.uber.org/zap"
)

type TemporalProxy struct {
	logger *zap.SugaredLogger

	initialized atomic.Bool
}

func NewTemporalProxy() service.Service {
	return &TemporalProxy{
		logger: zap.NewNop().Sugar(),
	}
}

func (tp *TemporalProxy) Name() string {
	return "TemporalProxy"
}

func (tp *TemporalProxy) Initialize(logger *zap.SugaredLogger) error {
	tp.logger = logger

	tp.initialized.Store(true)

	return nil
}

func (tp *TemporalProxy) Run() error {
	if !tp.initialized.Load() {
		return errors.New("unable to call .Run before setting initializing TemporalProxy")
	}

	return nil
}
