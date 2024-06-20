package usecase

import (
	"payment-system/domain/entity"
	interfaces "payment-system/pkg/interface"
)

type PaymentUsecase struct {
	PaymentRepo interfaces.IPaymentRepository
}

func (u *PaymentUsecase) ProcessBkm(initialization entity.InitializeBkmRequest) (entity.InitializeBkmResult, error) {
	response, auth, err := u.PaymentRepo.InitializeBkm(initialization)
	if err != nil {
		return entity.InitializeBkmResult{}, err
	}

	retrieveBkmResultDto := entity.RetrieveBkmResultRequest{
		Locale:         response.Data.Locale,
		ConversationID: response.Data.ConversationID,
		Token:          response.Data.Token,
	}

	return u.PaymentRepo.RetrieveBkmResult(retrieveBkmResultDto, auth)
}
func (u *PaymentUsecase) CreatePayment(createPaymentRequest entity.CreatePaymentRequest) (entity.CreatePaymentResult, error) {

	return u.PaymentRepo.CreatePayment(createPaymentRequest)
}
