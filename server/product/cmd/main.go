package main

import (
	"context"
	"github.com/omaqase/sato/product/internal/config"
	"github.com/omaqase/sato/product/internal/grpc"
	"github.com/omaqase/sato/product/internal/repository/postgres/category"
	"github.com/omaqase/sato/product/internal/repository/postgres/product"
	categoryService "github.com/omaqase/sato/product/internal/service/category"
	productService "github.com/omaqase/sato/product/internal/service/product"
	"github.com/omaqase/sato/product/pkg/pgx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			func() context.Context { return context.Background() },
			config.Load,
			config.ProvidePGXConfig,
			pgx.NewPGXPool,
			product.NewProductRepository,
			category.NewCategoryRepository,
			categoryService.NewCategoryService,
			productService.NewProductService,
			grpc.NewGRPCServer,
		),
		fx.Invoke(func(config *grpc.Server) {}),
	)

	app.Run()
}
