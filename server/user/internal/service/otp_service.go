package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/omaqase/sato/user/internal/repository/cache"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/notification"
	"github.com/omaqase/sato/user/pkg/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type IOTPService interface {
	StoreOTPCode(ctx context.Context) (*OTP, error)
	VerifyOTP(ctx context.Context, otp *OTP) (bool, error)
}

type OTPService struct {
	Cache              cache.ICache
	NotificationClient protobuf.NotificationServiceClient
}

func NewOTPService(cache cache.ICache) (IOTPService, error) {
	addr := fmt.Sprintf("%s:%s", "localhost", "5054")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	notificationClient := protobuf.NewNotificationServiceClient(conn)

	return &OTPService{
		Cache:              cache,
		NotificationClient: notificationClient,
	}, nil
}

type OTP struct {
	Token string
	Code  string
}

const CachePrefix = "otp"

func (s OTPService) StoreOTPCode(ctx context.Context) (*OTP, error) {
	code, err := crypto.GenerateOTPCode()
	if err != nil {
		return nil, err
	}

	otpToken := uuid.New()

	if err := s.Cache.Set(ctx, fmt.Sprintf("%s:%s", CachePrefix, otpToken.String()), code, time.Minute*5); err != nil {
		return nil, err
	}

	notificationRequest := &protobuf.SendNotificationRequest{
		Receiver:    "omaqase@gmail.com",
		Username:    "Sayat",
		Content:     code,
		SupportLink: "https://github.com/omaqase/sato",
	}

	_, err = s.NotificationClient.Send(ctx, notificationRequest)
	if err != nil {
		return nil, err
	}

	return &OTP{
		Token: otpToken.String(),
		Code:  code,
	}, nil
}

func (s OTPService) VerifyOTP(ctx context.Context, otp *OTP) (bool, error) {
	code, err := s.Cache.Get(ctx, fmt.Sprintf("%s:%s", CachePrefix, otp.Token))
	if err != nil {
		return false, err
	}
	log.Println(code)
	log.Println(otp.Code)

	if code != otp.Code {
		return false, nil
	}

	return true, nil
}
