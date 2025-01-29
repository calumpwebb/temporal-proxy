package main

import (
	"context"
	"temporal-proxy/pkg/proxy"
)

func main() {
	ctx := context.Background()

	tp := proxy.NewTemporalProxy("TEST-ID")
	tp.Run(ctx)
}
