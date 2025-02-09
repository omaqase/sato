package payment

type AddPaymentMethodRequest struct {
	UserID          string
	CardNumber      string
	Cardholder      string
	ExpirationMonth int
	ExpirationYear  int
	CVV             string
}

func NewAddPaymentMethodRequest(userID string, cardNumber string, cardholder string, expirationMonth int, expirationYear int, cvv string) *AddPaymentMethodRequest {
	return &AddPaymentMethodRequest{
		UserID:          userID,
		CardNumber:      cardNumber,
		Cardholder:      cardholder,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CVV:             cvv,
	}
}

type ProcessPaymentRequest struct {
	UserID          string
	PaymentMethodID string
	Amount          float64
	Currency        string
}

type PaymentMethod struct {
	UserID          string
	CardNumber      []byte
	CardholderName  string
	ExpirationMonth int
	ExpirationYear  int
	CVV             []byte
}

func NewProcessPaymentRequest(userID, paymentMethodID string, amount float64, currency string) *ProcessPaymentRequest {
	return &ProcessPaymentRequest{
		UserID:          userID,
		PaymentMethodID: paymentMethodID,
		Amount:          amount,
		Currency:        currency,
	}
}
