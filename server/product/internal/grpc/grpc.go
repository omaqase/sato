package grpc

import (
	"context"
	"fmt"
	"github.com/omaqase/sato/product/internal/config"
	protobuf "github.com/omaqase/sato/product/pkg/api/product/user"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Listener net.Listener
	Server   *grpc.Server
}

func StartGRPCServer(server *grpc.Server, listener net.Listener) error {
	go func() {
		if err := server.Serve(listener); err != nil {
			grpclog.Fatalf("failed to serve: %s", err)
		}
	}()

	grpclog.Infof("grpc server serving on %s", listener.Addr())
	return nil
}

func StopGRPCServer(server *grpc.Server) error {
	server.GracefulStop()
	return nil
}

func NewGRPCServer(lc fx.Lifecycle, categoryService protobuf.CategoryServiceServer, productService protobuf.ProductServiceServer, config config.Config) (*Server, error) {
	addr := fmt.Sprintf(":%s", config.App.Port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		grpclog.Fatalf("failed to listen tcp: %s", addr)
		return nil, err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	protobuf.RegisterProductServiceServer(server, productService)
	protobuf.RegisterCategoryServiceServer(server, categoryService)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return StartGRPCServer(server, listener)
		},
		OnStop: func(context.Context) error {
			return StopGRPCServer(server)
		},
	})

	return &Server{
		Listener: listener,
		Server:   server,
	}, nil
}
