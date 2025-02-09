package payment

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

type mockPool struct {
	pgxmock.PgxPoolIface
}

func TestCreatePaymentMethod(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := &Repository{pool: &mockPool{mock}}

	tests := []struct {
		name    string
		req     *AddPaymentMethodRequest
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			req: &AddPaymentMethodRequest{
				UserID:          "user-123",
				CardNumber:      "4242424242424242",
				Cardholder:      "John Doe",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				CVV:             "123",
			},
			mockFn: func() {
				mock.ExpectQuery("INSERT INTO payment_management.payment_methods").
					WithArgs(
						"user-123",
						pgxmock.AnyArg(),
						"John Doe",
						int16(12),
						int16(2025),
						pgxmock.AnyArg(),
					).
					WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(uuid.New()))
			},
			wantErr: false,
		},
		{
			name: "db_error",
			req: &AddPaymentMethodRequest{
				UserID:          "user-123",
				CardNumber:      "4242424242424242",
				Cardholder:      "John Doe",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				CVV:             "123",
			},
			mockFn: func() {
				mock.ExpectQuery("INSERT INTO payment_management.payment_methods").
					WillReturnError(assert.AnError)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			id, err := repo.CreatePaymentMethod(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, id)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, id)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestProcessPayment(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	repo := &Repository{pool: &mockPool{mock}}

	validPaymentMethodID := uuid.New()
	userID := "user-123"

	tests := []struct {
		name    string
		req     *ProcessPaymentRequest
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			req: &ProcessPaymentRequest{
				UserID:          userID,
				PaymentMethodID: validPaymentMethodID.String(),
				Amount:          100.00,
				Currency:        "USD",
			},
			mockFn: func() {
				mock.ExpectQuery("SELECT user_id, card_number_encrypted").
					WithArgs(validPaymentMethodID).
					WillReturnRows(pgxmock.NewRows([]string{
						"user_id", "card_number_encrypted", "cardholder_name",
						"expiration_month", "expiration_year", "cvv_encrypted",
					}).AddRow(
						userID, []byte("encrypted"), "John Doe",
						12, 2025, []byte("encrypted"),
					))

				mock.ExpectQuery("INSERT INTO payment_management.payments").
					WithArgs(
						userID,
						validPaymentMethodID.String(),
						100.00,
						"USD",
						"completed",
						pgxmock.AnyArg(),
					).
					WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(uuid.New()))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			id, err := repo.ProcessPayment(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, id)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, id)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
