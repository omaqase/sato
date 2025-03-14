package service

import (
	"errors"

	"github.com/google/uuid"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
)

//func ValidateSignUpRequest(request *protobuf.SignUpRequest) error {
//	if request == nil || len(request.Email) == 0 {
//		return errors.New("email is required")
//	}
//	if len(request.Password) < 8 {
//		return errors.New("password must be at least 8 characters long")
//	}
//	return nil
//}
//
//func ValidateSignInRequest(request *protobuf.SignInRequest) error {
//	if request == nil || len(request.Email) == 0 {
//		return errors.New("email is required")
//	}
//	if len(request.Password) < 8 {
//		return errors.New("password must be at least 8 characters long")
//	}
//	return nil
//}

//func ValidateVerifyOTPRequest(request *protobuf.VerifyOTPRequest) error {
//
//}

func ValidateGetSubscribedToPromotionUsersRequest(request *protobuf.GetSubscribedToPromotionUsersRequest) error {
	// Проверяем, что request не nil
	if request == nil {
		return errors.New("request cannot be nil")
	}

	// Проверяем Limit
	if request.Limit <= 0 || request.Limit > 1000 {
		return errors.New("limit must be in range 1-1000")
	}

	// Проверяем Offset
	if request.Offset < 0 || request.Offset > 1000 {
		return errors.New("offset must be in range 0-1000")
	}

	return nil
}

func ValidateUpdateUserRequest(request *protobuf.UpdateUserRequest) error {
	if request == nil {
		return errors.New("request is required")
	}

	return nil
}

func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid UUID format")
	}
	return nil
}

func ValidateGetFavoritesRequest(request *protobuf.GetFavoritesRequest) error {
	if request == nil {
		return errors.New("request cannot be nil")
	}

	if request.Limit <= 0 || request.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}

	if request.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	return nil
}
