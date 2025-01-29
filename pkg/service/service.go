package service

import "go.uber.org/zap"

type Service interface {
	Name() string
	Initialize(logger *zap.SugaredLogger) error
	Run() error
}
