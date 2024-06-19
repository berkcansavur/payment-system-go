package payment_interfaces

import "payment-system/domain/entity"

type IPaymentRepository interface {
	InitializeBkm(initialization entity.InitializeBkmRequest) (entity.InitializeBkmResult, entity.Authorization, error)
	RetrieveBkmResult(retrieveBkmResult entity.RetrieveBkmResultRequest, auth entity.Authorization) (entity.InitializeBkmResult, error)
	CreatePayment(createPayment entity.CreatePaymentRequest) (entity.CreatePaymentResult, error)
}
