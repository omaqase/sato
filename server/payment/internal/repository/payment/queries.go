package payment

// var CreatePaymentMethodSQL = `
// 	INSERT INTO payment_management.payment_methods (user_id, card_number_encrypted, cardholder_name, expiration_month, expiration_year, cvv_encrypted)
// 	VALUES ($1, $2, $3, $4, $5, $6)
// 	RETURNING id
// `

const ProcessPaymentSQL = `
    INSERT INTO payment_management.payments (
        user_id, 
        payment_method_id, 
        amount, 
        currency, 
        status, 
        transaction_reference
    )
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id
`

const GetPaymentMethodSQL = `
    SELECT user_id, card_number_encrypted, cardholder_name, expiration_month, expiration_year, cvv_encrypted
    FROM payment_management.payment_methods
    WHERE id = $1
`
