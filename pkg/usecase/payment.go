package usecase

import (
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"

	"github.com/berkcansavur/iyzico-authorization"
)

type PaymentUsecase struct {
	PaymentRepo interfaces.IPaymentRepository
}

func (u *PaymentUsecase) CreatePayment(createPaymentRequest iyzico.CreatePaymentRequest) (entity.CreatePaymentResult, error) {

	return u.PaymentRepo.CreatePayment(createPaymentRequest)
}
