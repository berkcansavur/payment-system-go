package usecase

import "payment-system/domain/entity"

type PaymentRepository interface {
    ProcessPayment(payment entity.Payment) (entity.PaymentResponse, error)
}

type PaymentUsecase struct {
    PaymentRepo PaymentRepository
}

func (u *PaymentUsecase) Execute(payment entity.Payment) (entity.PaymentResponse, error) {
    return u.PaymentRepo.ProcessPayment(payment)
}
