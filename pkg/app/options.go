package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Options struct {
	fxOptions fx.Option
}

type Option func(*Options) *Options

func NewDefaultOptions() *Options {
	return &Options{
		fxOptions: modules,
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(o *Options) *Options {
		o.fxOptions = fx.Options(
			o.fxOptions,
			fx.Replace(logger),
		)
		return o
	}
}
