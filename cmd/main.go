package main

import (
	"fmt"
	"temporal-proxy/pkg/app"

	"go.uber.org/zap"
)

func CustomLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func main() {
	logger, err := CustomLogger()
	if err != nil {
		panic(fmt.Sprintf("unable to set up logger: %e", err))
	}

	a := app.NewApp(
		app.WithLogger(logger),
	)
	a.Run()
}
