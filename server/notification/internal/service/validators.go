package service

import (
	"errors"
	protobuf "github.com/omaqase/sato/notification/pkg/api/v1/notification"
	"github.com/omaqase/sato/notification/pkg/utils"
)

func ValidateSendNotificationRequest(request *protobuf.SendNotificationRequest) error {
	if request == nil {
		return errors.New("request is null")
	}

	if len(request.Content) == 0 {
		return errors.New("content is empty")
	}

	if utils.ValidateEmail(request.Receiver) {
		return errors.New("receiver is invalid")
	}

	return nil
}
