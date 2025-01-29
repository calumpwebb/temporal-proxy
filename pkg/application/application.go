package application

import (
	"temporal-proxy/pkg/service"

	"go.uber.org/zap"
)

type Application struct {
	logger  *zap.SugaredLogger
	options *Options
}

func NewApplication(opts ...Option) *Application {
	// Set up all the options
	options := newDefaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	// Initialize all the services
	for _, service := range options.services {
		service.Initialize(options.logger)
	}

	return &Application{
		logger:  options.logger,
		options: options,
	}
}

func (a *Application) Run() error {
	a.logger.Infof("Running application with %d service(s)", len(a.options.services))

	// Run each service in it's own goroutine
	for _, service := range a.options.services {
		go a.runService(service)
	}

	return nil
}

func (a *Application) runService(service service.Service) error {
	a.logger.Infoln("Running service='%s'", service.Name())

	if err := service.Run(); err != nil {
		a.logger.Errorln("Error while running service='%s'. Error: %s", service.Name(), err)
		return err
	}

	return nil
}
