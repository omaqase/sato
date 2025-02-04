package main

import (
	"context"
	"github.com/omaqase/sato/user/internal/config"
	"github.com/omaqase/sato/user/internal/grpc"
	"github.com/omaqase/sato/user/internal/repository/cache"
	"github.com/omaqase/sato/user/internal/repository/postgres"
	"github.com/omaqase/sato/user/internal/repository/session"
	"github.com/omaqase/sato/user/internal/service"
	"github.com/omaqase/sato/user/pkg/pgx"
	"github.com/omaqase/sato/user/pkg/redis"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			func() context.Context { return context.Background() },
			config.Load,
			config.ProvidePGXConfig,
			config.ProvideRedisConfig,
			pgx.NewPGXPool,
			redis.NewRedisClient,
			cache.NewCache,
			service.NewOTPService,
			session.NewSessionStore,
			service.NewJWTService,
			postgres.NewUsersRepository,
			service.NewUsersService,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}
