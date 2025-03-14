package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/omaqase/sato/user/internal/config"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer  *grpc.Server
	userService protobuf.UserServiceServer
}

func NewGRPCServer(userService protobuf.UserServiceServer) *Server {
	grpcServer := grpc.NewServer()
	protobuf.RegisterUserServiceServer(grpcServer, userService)

	return &Server{
		grpcServer:  grpcServer,
		userService: userService,
	}
}

func NewListener(config config.Config) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.App.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %w", err)
	}
	return listener, nil
}

func StartGRPCServer(lc context.Context, server *Server, listener net.Listener) error {
	log.Printf("Starting gRPC server on %s", listener.Addr().String())
	return server.grpcServer.Serve(listener)
}
