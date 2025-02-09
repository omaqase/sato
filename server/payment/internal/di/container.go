package di

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaqase/sato/payment/internal/config"
	"github.com/omaqase/sato/payment/internal/repository/payment"
	ps "github.com/omaqase/sato/payment/internal/service/payment"
	pb "github.com/omaqase/sato/payment/pkg/api/v1/payment"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewContainer() *fx.App {
	return fx.New(
		fx.Provide(
			config.Load,
			NewDBPool,
			payment.NewRepository,
			ps.NewService,
			NewGRPCServer,
		),
		fx.Invoke(RegisterGRPCServices),
	)
}

func NewDBPool(cfg *config.Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.GetConnectionString())
}

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}

type GRPCParams struct {
	fx.In

	Config  *config.Config
	Server  *grpc.Server
	Service *ps.Service
}

func RegisterGRPCServices(lc fx.Lifecycle, p GRPCParams) {
	pb.RegisterPaymentServiceServer(p.Server, p.Service)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%s", p.Config.App.Port))
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			log.Printf("Starting gRPC server on port %s", p.Config.App.Port)

			go func() {
				if err := p.Server.Serve(listener); err != nil {
					log.Printf("Failed to serve: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Print("Stopping gRPC server")
			p.Server.GracefulStop()
			return nil
		},
	})
}
