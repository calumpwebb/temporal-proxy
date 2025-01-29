package main

import (
	"fmt"
	"os"
	"temporal-proxy/pkg/application"
	"temporal-proxy/pkg/temporalproxy"
	"time"

	"go.uber.org/zap"
)

var logger = zap.Must(
	zap.NewProduction(),
).Sugar()

func main() {
	/* Example Temporal Proxy setup */
	appTimeoutOptions := application.NewTimeoutOptions(
		application.WithShutdownTimeout(30*time.Second),
		application.WithForceShutdownTimeout(60*time.Second),
	)

	temporalProxyService := temporalproxy.NewTemporalProxy()

	app := application.NewApplication(
		application.WithLogger(logger),
		application.WithTimeoutOptions(appTimeoutOptions),
		application.WithService(temporalProxyService),
	)

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}
