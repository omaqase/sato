package service

import (
	"context"
	protobuf "github.com/omaqase/sato/notification/pkg/api/v1/notification"
	"github.com/resend/resend-go/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationService struct {
	client      *resend.Client
	userService IUserServiceClient
}

func NewNotificationService(client *resend.Client, userService IUserServiceClient) protobuf.NotificationServiceServer {
	return &NotificationService{client: client, userService: userService}
}

func (s NotificationService) Send(ctx context.Context, request *protobuf.SendNotificationRequest) (*protobuf.SendNotificationResponse, error) {
	if err := ValidateSendNotificationRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	finalHTML, err := GenerateFinalHTML(request.Content, request.Username, request.SupportLink)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate final HTML")
	}
	params := &resend.SendEmailRequest{
		From:    "satori@kiteo.app",
		To:      []string{request.Receiver},
		Subject: "Sato Notification",
		Html:    finalHTML,
	}
	sent, err := s.client.Emails.SendWithContext(ctx, params)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.SendNotificationResponse{
		Id: sent.Id,
	}, nil
}
