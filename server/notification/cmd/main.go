package main

import (
	"context"
	"github.com/omaqase/sato/notification/internal/config"
	"github.com/omaqase/sato/notification/internal/grpc"
	"github.com/omaqase/sato/notification/internal/service"
	"github.com/omaqase/sato/notification/pkg/resend"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			func() context.Context { return context.Background() },
			config.Load,
			resend.NewResendClient,
			service.NewUserServiceClient,
			service.NewNotificationService,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}
