package healthcheck

import "go.uber.org/fx"

var Module = fx.Module("healthcheck",
	fx.Provide(NewHealthCheckWorker),
	fx.Invoke(func(lifecycle fx.Lifecycle, worker *HealthCheckWorker) {
		lifecycle.Append(fx.Hook{
			OnStart: worker.Start,
			OnStop:  worker.Stop,
		})
	}),
)
