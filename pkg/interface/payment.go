package interfaces

import (
	"payment-system/domain/entity"

	"github.com/berkcansavur/iyzico-authorization"
)

type IPaymentRepository interface {
	InitializeBkm(initialization iyzico.InitializeBkmRequest) (entity.InitializeBkmResult, entity.Authorization, error)
	RetrieveBkmResult(retrieveBkmResult entity.RetrieveBkmResultRequest, auth entity.Authorization) (entity.InitializeBkmResult, error)
	CreatePayment(createPayment iyzico.CreatePaymentRequest) (entity.CreatePaymentResult, error)
}
