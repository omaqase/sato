package service

import (
	"context"
	"fmt"
	"github.com/omaqase/sato/notification/internal/config"
	protobuf "github.com/omaqase/sato/notification/pkg/api/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserServiceClient interface {
	GetSubscribedToPromotionUsers(ctx context.Context, in *protobuf.GetSubscribedToPromotionUsersRequest) (*protobuf.GetSubscribedToPromotionUsersResponse, error)
}

type UserServiceClient struct {
	Client protobuf.UserServiceClient
}

func NewUserServiceClient(config config.Config) (IUserServiceClient, error) {
	addr := fmt.Sprintf("%s:%s", config.UserService.Host, config.UserService.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := protobuf.NewUserServiceClient(conn)

	return &UserServiceClient{client}, nil
}

func (c UserServiceClient) GetSubscribedToPromotionUsers(ctx context.Context, in *protobuf.GetSubscribedToPromotionUsersRequest) (*protobuf.GetSubscribedToPromotionUsersResponse, error) {
	return c.Client.GetSubscribedToPromotionUsers(ctx, in)
}
