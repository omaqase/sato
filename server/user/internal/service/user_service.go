package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/omaqase/sato/user/internal/config"
	"github.com/omaqase/sato/user/internal/domain"
	"github.com/omaqase/sato/user/internal/repository/postgres"
	"github.com/omaqase/sato/user/internal/repository/session"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type UsersService struct {
	repository   postgres.IUsersRepository
	config       config.Config
	jwtService   IJWTService
	sessionStore session.IStore
	otpService   IOTPService
}

func (s UsersService) ParseJWT(ctx context.Context, request *protobuf.ParseJWTRequest) (*protobuf.ParseJWTResponse, error) {
	userID, err := s.jwtService.Parse(request.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &protobuf.ParseJWTResponse{
		UserId: userID,
	}, nil
}

func NewUsersService(repository postgres.IUsersRepository, config config.Config, jwtService IJWTService, sessionStore session.IStore, otpService IOTPService) protobuf.UserServiceServer {
	return &UsersService{
		repository:   repository,
		config:       config,
		jwtService:   jwtService,
		sessionStore: sessionStore,
		otpService:   otpService,
	}
}

func (s UsersService) SendOTP(ctx context.Context, request *protobuf.SendOTPRequest) (*protobuf.SendOTPResponse, error) {
	otp, err := s.otpService.StoreOTPCode(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &protobuf.SendOTPResponse{
		Token: otp.Token,
	}, nil
}

func (s UsersService) VerifyOTP(ctx context.Context, request *protobuf.VerifyOTPRequest) (*protobuf.VerifyOTPResponse, error) {
	otp := OTP{
		Token: request.Token,
		Code:  request.Code,
	}
	if err := s.verifyOTP(ctx, &otp); err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to verify OTP")
	}

	user, err := s.findOrCreateUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	return s.createSession(ctx, user.ID, user.Role)
}

func (s UsersService) verifyOTP(ctx context.Context, otp *OTP) error {
	_, err := s.otpService.VerifyOTP(ctx, otp)
	if err != nil {
		return fmt.Errorf("failed to verify OTP: %w", err)
	}
	return nil
}

func (s UsersService) findOrCreateUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	contract := &postgres.GetByEmailContract{
		Email: email,
	}
	user, err := s.repository.GetByEmail(ctx, contract)
	if err == nil {
		return user, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		createUserContract := &postgres.CreateContract{
			Email: email,
		}
		newUser, err := s.repository.Create(ctx, createUserContract)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to create user")
		}
		return newUser, nil
	}

	return nil, status.Error(codes.Internal, "failed to fetch user")
}

func (s UsersService) createSession(ctx context.Context, userID, role string) (*protobuf.VerifyOTPResponse, error) {
	log.Println(role)
	accessToken, err := s.jwtService.NewJWT(userID, role, s.config.Security.JWTExpiration)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate access token")
	}
	refreshToken, err := s.jwtService.NewRefreshToken()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate refresh token")
	}
	session := domain.Session{
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(s.config.Security.JWTExpiration),
		Role:         role,
	}
	tokens := &protobuf.VerifyOTPResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsNewUser:    true,
	}
	if err := s.sessionStore.SetSession(ctx, userID, session, s.config.Security.JWTExpiration); err != nil {
		return nil, status.Error(codes.Internal, "failed to store session")
	}
	return tokens, nil
}

func (s UsersService) LogOut(ctx context.Context, request *protobuf.LogOutRequest) (*protobuf.LogOutResponse, error) {
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "missing user ID in metadata")
	}
	if err := s.sessionStore.DeleteSession(ctx, userID); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete session")
	}
	return &protobuf.LogOutResponse{}, nil
}

func (s UsersService) RefreshToken(ctx context.Context, request *protobuf.RefreshTokenRequest) (*protobuf.RefreshTokenResponse, error) {
	if request == nil || request.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "missing user ID in metadata")
	}
	session, err := s.sessionStore.GetSessionByRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired refresh token")
		}
		return nil, status.Error(codes.Internal, "failed to retrieve session")
	}
	newAccessToken, err := s.jwtService.NewJWT(userID, session.Role, s.config.Security.JWTExpiration)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate new access token")
	}
	if err := s.sessionStore.UpdateSession(ctx, userID, *session, s.config.Security.RefreshTokenExpiration); err != nil {
		return nil, status.Error(codes.Internal, "failed to update session")
	}
	return &protobuf.RefreshTokenResponse{AccessToken: newAccessToken}, nil
}

func (s UsersService) Update(ctx context.Context, request *protobuf.UpdateUserRequest) (*protobuf.UpdateUserResponse, error) {
	if err := ValidateUpdateUserRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "missing user ID in metadata")
	}
	contract := &postgres.UpdateContract{
		ID:         userID,
		FirstName:  *request.FirstName,
		LastName:   *request.LastName,
		Phone:      *request.Phone,
		Promotions: *request.Promotions,
	}
	user, err := s.repository.Update(ctx, contract)
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, postgres.ErrInvalidSignUpCredentials) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, postgres.ErrValidationFailed) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UpdateUserResponse{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Phone:      &user.Phone,
		Promotions: &user.Promotions,
		UpdatedAt:  user.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (s UsersService) GetSubscribedToPromotionUsers(ctx context.Context, request *protobuf.GetSubscribedToPromotionUsersRequest) (*protobuf.GetSubscribedToPromotionUsersResponse, error) {
	if err := ValidateGetSubscribedToPromotionUsersRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	users, err := s.repository.GetSubscribedToPromotion(ctx, int(request.Limit), int(request.Offset))
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to retrieve subscribed users")
	}
	responseItems := make([]*protobuf.GetSubscribedToPromotionUsersResponseItem, len(users))
	for i, user := range users {
		responseItems[i] = &protobuf.GetSubscribedToPromotionUsersResponseItem{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}
	return &protobuf.GetSubscribedToPromotionUsersResponse{Emails: responseItems}, nil
}
