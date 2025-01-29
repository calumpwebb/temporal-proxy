package application

import "time"

const (
	DefaultShutdownTimeout      = 30 * time.Second
	DefaultForceShutdownTimeout = 60 * time.Second
)

type TimeoutOptions struct {
	Shutdown      time.Duration
	ForceShutdown time.Duration
}

type TimeoutOption func(*TimeoutOptions)

func newDefaultTimeoutOptions() *TimeoutOptions {
	return &TimeoutOptions{
		Shutdown:      DefaultShutdownTimeout,
		ForceShutdown: DefaultForceShutdownTimeout,
	}
}

func NewTimeoutOptions(opts ...TimeoutOption) *TimeoutOptions {
	options := newDefaultTimeoutOptions()
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithShutdownTimeout(duration time.Duration) TimeoutOption {
	return func(to *TimeoutOptions) {
		to.Shutdown = duration
	}
}

func WithForceShutdownTimeout(duration time.Duration) TimeoutOption {
	return func(to *TimeoutOptions) {
		to.ForceShutdown = duration
	}
}
