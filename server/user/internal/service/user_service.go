package service

import (
	"context"
	"errors"
	"github.com/omaqase/sato/user/internal/repository/postgres"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UsersService struct {
	repository postgres.IUsersRepository

	authService IAuthService
}

func NewUsersService(repository postgres.IUsersRepository, authService IAuthService) protobuf.UserServiceServer {
	return &UsersService{repository: repository, authService: authService}
}

func (s UsersService) SignUp(ctx context.Context, request *protobuf.SignUpRequest) (*protobuf.SignUpResponse, error) {
	if err := ValidateSignUpRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := postgres.NewCreateUserContract(request.Email, request.Password, request.FirstName, request.LastName)

	user, err := s.repository.Create(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrInvalidSignUpCredentials) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.SignUpResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s UsersService) SignIn(ctx context.Context, request *protobuf.SignInRequest) (*protobuf.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UsersService) LogOut(ctx context.Context, request *protobuf.LogOutRequest) (*protobuf.LogOutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UsersService) RefreshToken(ctx context.Context, request *protobuf.RefreshTokenRequest) (*protobuf.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UsersService) Update(ctx context.Context, response *protobuf.UpdateUserRequest) (*protobuf.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UsersService) GetSubscribedToPromotionUsers(ctx context.Context, request *protobuf.GetSubscribedToPromotionUsersRequest) (*protobuf.GetSubscribedToPromotionUsersResponse, error) {
	if err := ValidateGetSubscribedToPromotionUsersRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := postgres.NewGetSubscribedToPromotionContract(int(request.Limit), request.Cursor)

	emails, err := s.repository.GetSubscribedToPromotions(ctx, contract)
	if err != nil {
		return nil, err
	}

	var responseItems []*protobuf.GetSubscribedToPromotionUsersResponseItem
	for _, email := range *emails {
		responseItems = append(responseItems, &protobuf.GetSubscribedToPromotionUsersResponseItem{
			Email: email,
		})
	}

	return &protobuf.GetSubscribedToPromotionUsersResponse{
		Emails: responseItems,
	}, nil
}
