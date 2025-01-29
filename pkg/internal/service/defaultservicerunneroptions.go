package service

import "time"

const DefaultShutdownTimeout = 30 * time.Second

type DefaultServiceRunnerOptions struct {
	ShutdownTimeout time.Duration
}

type DefaultServiceRunnerOption func(*DefaultServiceRunnerOptions)

func NewDefaultServiceRunnerOptions(opts ...DefaultServiceRunnerOption) DefaultServiceRunnerOptions {
	options := DefaultServiceRunnerOptions{
		ShutdownTimeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return options
}

func WithShutdownTimeout(d time.Duration) DefaultServiceRunnerOption {
	return func(o *DefaultServiceRunnerOptions) {
		o.ShutdownTimeout = d
	}
}
