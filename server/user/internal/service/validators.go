package service

import (
	"errors"
	"github.com/google/uuid"
	protobuf "github.com/omaqase/sato/user/pkg/api/v1/user"
)

func ValidateSignUpRequest(request *protobuf.SignUpRequest) error {
	if len(request.Email) == 0 {
		return errors.New("invalid signup credentials")
	}

	return nil
}

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
