package payment_interfaces

import "payment-system/domain/entity"

type IPaymentRepository interface {
	InitializeBkm(initialization entity.InitializeBkmRequest) (entity.PaymentResponse, entity.Authorization, error)
	RetrieveBkmResult(retrieveBkmResult entity.RetrieveBkmResultRequest, auth entity.Authorization) (entity.PaymentResponse, error)
}
