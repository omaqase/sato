package payment

import (
	"context"
	"testing"

	"github.com/omaqase/sato/payment/internal/repository/payment"
	pb "github.com/omaqase/sato/payment/pkg/api/v1/payment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePaymentMethod(ctx context.Context, req *payment.AddPaymentMethodRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) ProcessPayment(ctx context.Context, req *payment.ProcessPaymentRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) GetPaymentMethod(ctx context.Context, id string) (*payment.PaymentMethod, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*payment.PaymentMethod), args.Error(1)
}

func TestAddPaymentMethod(t *testing.T) {
	tests := []struct {
		name    string
		req     *pb.AddPaymentMethodRequest
		mockFn  func(*MockRepository)
		wantErr bool
	}{
		{
			name: "success",
			req: &pb.AddPaymentMethodRequest{
				UserId:          "user-123",
				CardNumber:      "4242424242424242",
				CardholderName:  "John Doe",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				Cvv:             "123",
			},
			mockFn: func(m *MockRepository) {
				m.On("CreatePaymentMethod", mock.Anything, mock.MatchedBy(func(req *payment.AddPaymentMethodRequest) bool {
					return req.UserID == "user-123"
				})).Return("pm-123", nil)
			},
			wantErr: false,
		},
		{
			name: "invalid_request",
			req: &pb.AddPaymentMethodRequest{
				UserId: "",
			},
			mockFn:  func(m *MockRepository) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockFn(mockRepo)

			s := NewService(mockRepo)
			resp, err := s.AddPaymentMethod(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.PaymentMethodId)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
