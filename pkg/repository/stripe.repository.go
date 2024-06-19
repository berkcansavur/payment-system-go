package repository

import (
	"payment-system/domain/entity"
)

type StripeRepository struct{}

func (r *StripeRepository) ProcessPayment(payment entity.Payment) (entity.PaymentResult, error) {
	return entity.PaymentResult{Status: "success", Message: "Stripe payment successful"}, nil
}
