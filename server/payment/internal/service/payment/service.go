package payment

import (
	"context"

	"github.com/omaqase/sato/payment/internal/repository/payment"
	pb "github.com/omaqase/sato/payment/pkg/api/v1/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedPaymentServiceServer
	repo payment.IRepository
}

func NewService(repo payment.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddPaymentMethod(ctx context.Context, req *pb.AddPaymentMethodRequest) (*pb.AddPaymentMethodResponse, error) {
	if err := validateAddPaymentMethod(req); err != nil {
		return nil, err
	}

	contract := payment.NewAddPaymentMethodRequest(
		req.UserId,
		req.CardNumber,
		req.CardholderName,
		int(req.ExpirationMonth),
		int(req.ExpirationYear),
		req.Cvv,
	)

	id, err := s.repo.CreatePaymentMethod(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create payment method")
	}

	return &pb.AddPaymentMethodResponse{
		PaymentMethodId: id,
	}, nil
}

func validateAddPaymentMethod(req *pb.AddPaymentMethodRequest) error {
	if req.UserId == "" {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if req.CardNumber == "" {
		return status.Error(codes.InvalidArgument, "card_number is required")
	}
	if req.CardholderName == "" {
		return status.Error(codes.InvalidArgument, "cardholder_name is required")
	}
	if req.ExpirationMonth < 1 || req.ExpirationMonth > 12 {
		return status.Error(codes.InvalidArgument, "invalid expiration_month")
	}
	if req.ExpirationYear < 2024 {
		return status.Error(codes.InvalidArgument, "invalid expiration_year")
	}
	if req.Cvv == "" {
		return status.Error(codes.InvalidArgument, "cvv is required")
	}
	return nil
}

func (s *Service) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	if err := validateProcessPayment(req); err != nil {
		return nil, err
	}

	contract := payment.NewProcessPaymentRequest(
		req.UserId,
		req.PaymentMethodId,
		req.Amount,
		req.Currency,
	)

	paymentID, err := s.repo.ProcessPayment(ctx, contract)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to process payment")
	}

	return &pb.ProcessPaymentResponse{
		TransactionId: paymentID,
		Status:        "completed",
	}, nil
}

func validateProcessPayment(req *pb.ProcessPaymentRequest) error {
	if req.UserId == "" {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if req.PaymentMethodId == "" {
		return status.Error(codes.InvalidArgument, "payment_method_id is required")
	}
	if req.Amount <= 0 {
		return status.Error(codes.InvalidArgument, "amount must be greater than 0")
	}
	if req.Currency == "" {
		return status.Error(codes.InvalidArgument, "currency is required")
	}
	return nil
}
