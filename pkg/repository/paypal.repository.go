package repository

import (
	"payment-system/domain/entity"
)

type PayPalRepository struct{}

func (r *PayPalRepository) ProcessPayment(payment entity.Payment) (entity.PaymentResult, error) {
	return entity.PaymentResult{Status: "success", Message: "PayPal payment successful"}, nil
}
