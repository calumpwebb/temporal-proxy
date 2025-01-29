package application

import (
	"temporal-proxy/pkg/service"

	"go.uber.org/zap"
)

var DefaultZapLogger = zap.Must(zap.NewDevelopment()).Sugar()

type Options struct {
	logger   *zap.SugaredLogger
	timeouts *TimeoutOptions
	services []service.Service
}

type Option func(*Options)

func newDefaultOptions() *Options {
	return &Options{
		logger:   DefaultZapLogger,
		timeouts: newDefaultTimeoutOptions(),
		services: nil,
	}
}

func WithLogger(logger *zap.SugaredLogger) Option {
	return func(o *Options) {
		o.logger = logger
	}
}

func WithService(service service.Service) Option {
	return func(o *Options) {
		o.services = append(o.services, service)
	}
}

func WithTimeoutOptions(to *TimeoutOptions) Option {
	return func(o *Options) {
		o.timeouts = to
	}
}
