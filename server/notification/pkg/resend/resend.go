package resend

import (
	"errors"
	"github.com/omaqase/sato/notification/internal/config"
	"github.com/resend/resend-go/v2"
)

func NewResendClient(config config.Config) (*resend.Client, error) {
	client := resend.NewClient(config.Resend.APIKey)

	if client == nil {
		return nil, errors.New("creating resend api client failed")
	}

	return client, nil
}
