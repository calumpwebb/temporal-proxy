package temporalproxy

import "go.uber.org/fx"

var Module = fx.Module("temporalproxy",
	fx.Provide(NewTemporalProxyWorker),
	fx.Invoke(func(lifecycle fx.Lifecycle, worker *TemporalProxyWorker) {
		lifecycle.Append(fx.Hook{
			OnStart: worker.Start,
			OnStop:  worker.Stop,
		})
	}),
)
