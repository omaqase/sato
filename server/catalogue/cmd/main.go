package main

import (
	"context"

	"github.com/omaqase/sato/catalogue/internal/config"
	"github.com/omaqase/sato/catalogue/internal/grpc"
	categoryRepository "github.com/omaqase/sato/catalogue/internal/repository/category"
	productRepository "github.com/omaqase/sato/catalogue/internal/repository/product"
	"github.com/omaqase/sato/catalogue/internal/service/category"
	"github.com/omaqase/sato/catalogue/internal/service/jwt"
	"github.com/omaqase/sato/catalogue/internal/service/product"
	userService "github.com/omaqase/sato/catalogue/internal/service/user"
	"github.com/omaqase/sato/catalogue/pkg/pgx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			provideContext,
			config.Load,
			config.ProvidePGXConfig,
			pgx.NewPGXPool,
			productRepository.NewRepository,
			categoryRepository.NewRepository,
			userService.NewClient,
			jwt.NewService,
			category.NewService,
			product.NewService,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}

func provideContext() context.Context {
	return context.Background()
}
