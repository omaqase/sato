package service

import (
	"context"
	"errors"
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
		OtpToken: otp.Token,
	}, nil
}

func (s UsersService) VerifyOTP(ctx context.Context, request *protobuf.VerifyOTPRequest) (*protobuf.VerifyOTPResponse, error) {
	//md, _ := metadata.FromIncomingContext(ctx)
	//log.Println(md)
	otp := OTP{
		Token: request.OtpToken,
		Code:  request.Code,
	}

	valid, err := s.otpService.VerifyOTP(ctx, &otp)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	log.Println(valid)

	contract := &postgres.GetByEmailContract{
		Email: request.Email,
	}

	user, err := s.repository.GetByEmail(ctx, contract)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createUserContract := &postgres.CreateContract{
				Email:     request.Email,
				FirstName: "qwe",
				LastName:  "qwe",
			}

			us, err := s.repository.Create(ctx, createUserContract)
			if err != nil {
				return nil, status.Error(codes.Unknown, err.Error())
			}

			log.Println(us)

			return s.createSession(ctx, us.ID)
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}

	log.Println(user)

	return s.createSession(ctx, user.ID)
}

func (s UsersService) GetUnverifiedEmailUsers(ctx context.Context, request *protobuf.GetUnverifiedEmailUsersRequest) (*protobuf.GetUnverifiedEmailUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UsersService) createSession(ctx context.Context, userID string) (*protobuf.VerifyOTPResponse, error) {
	accessToken, err := s.jwtService.NewJWT(userID, s.config.Security.JWTExpiration)
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
	newAccessToken, err := s.jwtService.NewJWT(userID, s.config.Security.JWTExpiration)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate new access token")
	}
	if err := s.sessionStore.UpdateSession(ctx, userID, *session, s.config.Security.RefreshTokenExpiration); err != nil {
		return nil, status.Error(codes.Internal, "failed to update session")
	}
	return &protobuf.RefreshTokenResponse{AccessToken: newAccessToken}, nil
}

func (s UsersService) Update(ctx context.Context, request *protobuf.UpdateUserRequest) (*protobuf.UpdateUserResponse, error) {
	//if err := ValidateUpdateUserRequest(request); err != nil {
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}
	//userID, exists := s.ExtractUserIDFromMetadata(ctx)
	//if !exists {
	//	return nil, status.Error(codes.InvalidArgument, "missing user ID in metadata")
	//}
	//contract := postgres.NewUpdateUserContract(userID, request.FirstName, request.LastName, request.Phone, request.Promotions)
	//user, err := s.repository.Update(ctx, contract)
	//if err != nil {
	//	if errors.Is(err, postgres.ErrUserNotFound) {
	//		return nil, status.Error(codes.NotFound, err.Error())
	//	}
	//	if errors.Is(err, postgres.ErrInvalidSignUpCredentials) {
	//		return nil, status.Error(codes.InvalidArgument, err.Error())
	//	}
	//	if errors.Is(err, postgres.ErrValidationFailed) {
	//		return nil, status.Error(codes.InvalidArgument, err.Error())
	//	}
	//	return nil, status.Error(codes.Internal, err.Error())
	//}
	//
	//return &protobuf.UpdateUserResponse{
	//	FirstName:  user.FirstName,
	//	LastName:   user.LastName,
	//	Phone:      &user.Phone,
	//	Promotions: &user.IsSubscribedToPromotions,
	//	UpdatedAt:  user.UpdatedAt.Format(time.DateTime),
	//}, nil
	return nil, nil
}

func (s UsersService) GetSubscribedToPromotionUsers(ctx context.Context, request *protobuf.GetSubscribedToPromotionUsersRequest) (*protobuf.GetSubscribedToPromotionUsersResponse, error) {
	//if err := ValidateGetSubscribedToPromotionUsersRequest(request); err != nil {
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}
	//contract := postgres.NewGetSubscribedToPromotionContract(int(request.Limit), request.Cursor)
	//emails, err := s.repository.GetSubscribedToPromotions(ctx, contract)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, "failed to retrieve subscribed users")
	//}
	//responseItems := make([]*protobuf.GetSubscribedToPromotionUsersResponseItem, len(emails))
	//for i, email := range emails {
	//	responseItems[i] = &protobuf.GetSubscribedToPromotionUsersResponseItem{Email: email}
	//}
	//return &protobuf.GetSubscribedToPromotionUsersResponse{Emails: responseItems}, nil
	return nil, nil
}
