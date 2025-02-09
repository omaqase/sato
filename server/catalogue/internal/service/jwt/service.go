package jwt

import (
	"context"
	"github.com/omaqase/sato/catalogue/internal/service/user"
	"github.com/omaqase/sato/catalogue/pkg/api/v1/user"
	"github.com/omaqase/sato/catalogue/pkg/utils"
)

type IService interface {
	Parse(ctx context.Context) (string, error)
}

type Service struct {
	client user.IClient
}

func NewService(client user.IClient) IService {
	return &Service{client: client}
}

func (s Service) Parse(ctx context.Context) (string, error) {
	tokenMD, err := utils.ParseMetadata(ctx)
	if err != nil {
		return "", err
	}

	in := &protobuf.ParseJWTRequest{
		AccessToken: tokenMD,
	}

	token, err := s.client.ParseJWT(ctx, in)
	if err != nil {
		return "", err
	}

	return token.UserId, nil
}
