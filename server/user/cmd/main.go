package main

import (
	"context"
	"github.com/omaqase/sato/user/internal/config"
	"github.com/omaqase/sato/user/internal/grpc"
	"github.com/omaqase/sato/user/internal/repository/postgres"
	"github.com/omaqase/sato/user/internal/service"
	"github.com/omaqase/sato/user/pkg/pgx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			func() context.Context {
				return context.Background()
			},
			config.Load,
			config.ProvidePGXConfig,
			pgx.NewPGXPool,
			postgres.NewUsersRepository,
			service.NewUsersService,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}
