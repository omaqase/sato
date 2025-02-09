package payment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	CreatePaymentMethod(ctx context.Context, req *AddPaymentMethodRequest) (string, error)
	ProcessPayment(ctx context.Context, req *ProcessPaymentRequest) (string, error)
	GetPaymentMethod(ctx context.Context, id string) (*PaymentMethod, error)
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) IRepository {
	return &Repository{pool: pool}
}

var CreatePaymentMethodSQL = `
	INSERT INTO payment_management.payment_methods (user_id, card_number_encrypted, cardholder_name, expiration_month, expiration_year, cvv_encrypted)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
`

func (r *Repository) CreatePaymentMethod(ctx context.Context, req *AddPaymentMethodRequest) (string, error) {
	var id uuid.UUID

	err := r.pool.QueryRow(ctx, CreatePaymentMethodSQL,
		req.UserID,
		EncryptCardNumber(req.CardNumber),
		req.Cardholder,
		req.ExpirationMonth,
		req.ExpirationYear,
		EncryptCVV(req.CVV),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *Repository) ProcessPayment(ctx context.Context, req *ProcessPaymentRequest) (string, error) {
	paymentMethod, err := r.GetPaymentMethod(ctx, req.PaymentMethodID)
	if err != nil {
		return "", fmt.Errorf("failed to get payment method: %w", err)
	}

	if paymentMethod.UserID != req.UserID {
		return "", fmt.Errorf("payment method does not belong to user")
	}

	// В реальном приложении здесь был бы вызов платежного шлюза
	transactionRef := uuid.New().String()

	var paymentID uuid.UUID
	err = r.pool.QueryRow(ctx, ProcessPaymentSQL,
		req.UserID,
		req.PaymentMethodID,
		req.Amount,
		req.Currency,
		"completed", // В реальном приложении статус зависел бы от ответа платежного шлюза
		transactionRef,
	).Scan(&paymentID)

	if err != nil {
		return "", fmt.Errorf("failed to save payment: %w", err)
	}

	return paymentID.String(), nil
}

func (r *Repository) GetPaymentMethod(ctx context.Context, id string) (*PaymentMethod, error) {
	paymentMethodID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid payment method id: %w", err)
	}

	var pm PaymentMethod
	err = r.pool.QueryRow(ctx, GetPaymentMethodSQL, paymentMethodID).Scan(
		&pm.UserID,
		&pm.CardNumber,
		&pm.CardholderName,
		&pm.ExpirationMonth,
		&pm.ExpirationYear,
		&pm.CVV,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get payment method: %w", err)
	}

	return &pm, nil
}
