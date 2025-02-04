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
	if request.Limit <= 0 || request.Limit > 1000 {
		return errors.New("limit must be in range 1-1000")
	}
	if request.Cursor != "" {
		if _, err := uuid.Parse(request.Cursor); err != nil {
			return errors.New("invalid cursor format")
		}
	}
	return nil
}

func ValidateUpdateUserRequest(request *protobuf.UpdateUserRequest) error {
	if request == nil {
		return errors.New("request is required")
	}

	return nil
}
