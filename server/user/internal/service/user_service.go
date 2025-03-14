package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/omaqase/sato/user/internal/config"
	"github.com/omaqase/sato/user/internal/domain"
	"github.com/omaqase/sato/user/internal/repository/postgres"
	"github.com/omaqase/sato/user/internal/repository/session"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersService struct {
	repository   postgres.IUsersRepository
	config       config.Config
	jwtService   IJWTService
	sessionStore session.IStore
	otpService   IOTPService
	cartService  domain.ICartService
}

func (s UsersService) MakeSeller(ctx context.Context, request *protobuf.MakeSellerRequest) (*protobuf.MakeSellerResponse, error) {
	if err := ValidateUUID(request.UserId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := s.repository.GetById(ctx, request.UserId); err != nil {
		log.Println("there")
		return nil, status.Error(codes.InvalidArgument, "no user with such id")
	}

	_, err := s.repository.MakeSeller(ctx, request.UserId)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "no user with such id")
	}

	out := &protobuf.MakeSellerResponse{
		Status: 0,
	}

	return out, nil
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

func NewUsersService(repository postgres.IUsersRepository, config config.Config, jwtService IJWTService, sessionStore session.IStore, otpService IOTPService, cartService domain.ICartService) protobuf.UserServiceServer {
	return &UsersService{
		repository:   repository,
		config:       config,
		jwtService:   jwtService,
		sessionStore: sessionStore,
		otpService:   otpService,
		cartService:  cartService,
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

	user, isNewUser, err := s.findOrCreateUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtService.NewJWT(user.ID, user.Role, s.config.Security.JWTExpiration)
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
		Role:         user.Role,
	}

	if err := s.sessionStore.SetSession(ctx, user.ID, session, s.config.Security.JWTExpiration); err != nil {
		return nil, status.Error(codes.Internal, "failed to store session")
	}

	response := &protobuf.VerifyOTPResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsNewUser:    isNewUser,
	}

	if !isNewUser {
		log.Println("user", user)
		response.User = &protobuf.UserDTO{
			Email:      user.Email,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Phone:      user.Phone,
			Promotions: &user.Promotions,
		}
	}

	return response, nil
}

func (s UsersService) verifyOTP(ctx context.Context, otp *OTP) error {
	_, err := s.otpService.VerifyOTP(ctx, otp)
	if err != nil {
		return fmt.Errorf("failed to verify OTP: %w", err)
	}
	return nil
}

func (s UsersService) findOrCreateUserByEmail(ctx context.Context, email string) (*domain.User, bool, error) {
	contract := &postgres.GetByEmailContract{
		Email: email,
	}
	user, err := s.repository.GetByEmail(ctx, contract)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createUserContract := &postgres.CreateContract{
				Email: email,
			}
			newUser, err := s.repository.Create(ctx, createUserContract)
			if err != nil {
				return nil, false, status.Error(codes.Internal, fmt.Sprintf("failed to create user: %v", err))
			}
			return newUser, true, nil
		}
		return nil, false, status.Error(codes.Internal, fmt.Sprintf("failed to fetch user: %v", err))
	}

	return user, false, nil
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
		Phone:      user.Phone,
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

	return &protobuf.GetSubscribedToPromotionUsersResponse{
		Emails: responseItems,
	}, nil
}

func (s UsersService) AddToFavorites(ctx context.Context, request *protobuf.AddToFavoritesRequest) (*protobuf.AddToFavoritesResponse, error) {
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	if err := ValidateUUID(request.ProductId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	contract := &postgres.AddToFavoritesContract{
		UserID:    userID,
		ProductID: request.ProductId,
	}

	favorite, err := s.repository.AddToFavorites(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to add to favorites")
	}

	return &protobuf.AddToFavoritesResponse{
		Id:        favorite.ID,
		ProductId: favorite.ProductID,
		CreatedAt: favorite.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s UsersService) RemoveFromFavorites(ctx context.Context, request *protobuf.RemoveFromFavoritesRequest) (*protobuf.RemoveFromFavoritesResponse, error) {
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	if err := ValidateUUID(request.ProductId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	err := s.repository.RemoveFromFavorites(ctx, userID, request.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to remove from favorites")
	}

	return &protobuf.RemoveFromFavoritesResponse{}, nil
}

func (s UsersService) GetFavorites(ctx context.Context, request *protobuf.GetFavoritesRequest) (*protobuf.GetFavoritesResponse, error) {
	userID, exists := s.ExtractUserIDFromMetadata(ctx)
	if !exists {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	if err := ValidateGetFavoritesRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	contract := &postgres.GetFavoritesContract{
		UserID: userID,
		Limit:  int(request.Limit),
		Offset: int(request.Offset),
	}

	favorites, total, err := s.repository.GetFavorites(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get favorites")
	}

	items := make([]*protobuf.FavoriteItem, len(favorites))
	for i, favorite := range favorites {
		items[i] = &protobuf.FavoriteItem{
			Id:        favorite.ID,
			ProductId: favorite.ProductID,
			CreatedAt: favorite.CreatedAt.Format(time.RFC3339),
		}
	}

	return &protobuf.GetFavoritesResponse{
		Items: items,
		Total: int32(total),
	}, nil
}

func (s *UsersService) AddToCart(ctx context.Context, req *protobuf.AddToCartRequest) (*protobuf.AddToCartResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	err = s.cartService.AddToCart(ctx, domain.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  int(req.Quantity),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to add to cart: %v", err))
	}

	return &protobuf.AddToCartResponse{}, nil
}

func (s *UsersService) GetCart(ctx context.Context, req *protobuf.GetCartRequest) (*protobuf.GetCartResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	items, err := s.cartService.GetCart(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get cart: %v", err))
	}

	pbItems := make([]*protobuf.CartItem, len(items))
	for i, item := range items {
		pbItems[i] = &protobuf.CartItem{
			Id:        item.ID.String(),
			ProductId: item.ProductID.String(),
			Quantity:  int32(item.Quantity),
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
		}
	}

	return &protobuf.GetCartResponse{
		Items: pbItems,
		Total: int32(len(items)),
	}, nil
}

func (s *UsersService) UpdateCartItem(ctx context.Context, req *protobuf.UpdateCartItemRequest) (*protobuf.UpdateCartItemResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	err = s.cartService.UpdateCartItem(ctx, userID, productID, int(req.Quantity))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update cart item: %v", err))
	}

	return &protobuf.UpdateCartItemResponse{}, nil
}

func (s *UsersService) RemoveFromCart(ctx context.Context, req *protobuf.RemoveFromCartRequest) (*protobuf.RemoveFromCartResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}

	err = s.cartService.RemoveFromCart(ctx, userID, productID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to remove from cart: %v", err))
	}

	return &protobuf.RemoveFromCartResponse{}, nil
}

func (s *UsersService) ClearCart(ctx context.Context, req *protobuf.ClearCartRequest) (*protobuf.ClearCartResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = s.cartService.ClearCart(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to clear cart: %v", err))
	}

	return &protobuf.ClearCartResponse{}, nil
}

func getUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID not found in context")
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	return id, nil
}
