package usecase

import (
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type PaymentUsecase struct {
	PaymentRepo interfaces.IPaymentRepository
}

func (u *PaymentUsecase) CreatePayment(createPaymentRequest entity.CreatePaymentRequest) (entity.CreatePaymentResult, error) {

	return u.PaymentRepo.CreatePayment(createPaymentRequest)
}
