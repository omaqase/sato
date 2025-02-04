package main

import (
	"context"
	"github.com/omaqase/sato/gateway/config"
	"github.com/omaqase/sato/gateway/internal/proxy"
	"log"
)

func main() {
	configs, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	proxy := proxy.NewProxy(configs)
	if err := proxy.RegisterServices(ctx); err != nil {
		log.Fatalf("Failed to register services: %v", err)
	}

	if err := proxy.Serve(); err != nil {
		log.Fatalf("Failed to start proxy server: %v", err)
	}
}
