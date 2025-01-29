package app

import (
	"go.uber.org/fx"
)

func NewApp(options ...Option) *fx.App {
	defaultOptions := NewDefaultOptions()

	for _, opt := range options {
		defaultOptions = opt(defaultOptions) // Apply options correctly
	}

	return fx.New(
		defaultOptions.fxOptions,
	)
}
