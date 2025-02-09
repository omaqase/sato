package user

import (
	"context"
	"fmt"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClient interface {
	ParseJWT(ctx context.Context, in *protobuf.ParseJWTRequest) (*protobuf.ParseJWTResponse, error)
}

type Client struct {
	client protobuf.UserServiceClient
}

func NewClient() (IClient, error) {
	addr := fmt.Sprintf("%s:%s", "localhost", "5053")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := protobuf.NewUserServiceClient(conn)

	return &Client{client: client}, nil
}

func (c Client) ParseJWT(ctx context.Context, in *protobuf.ParseJWTRequest) (*protobuf.ParseJWTResponse, error) {
	return c.client.ParseJWT(ctx, in)
}
