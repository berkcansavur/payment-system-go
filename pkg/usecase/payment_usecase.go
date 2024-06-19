package usecase

import (
	"payment-system/domain/entity"
	payment_interfaces "payment-system/pkg/interface"
)

type PaymentUsecase struct {
	PaymentRepo payment_interfaces.IPaymentRepository
}

func (u *PaymentUsecase) ProcessBkm(initialization entity.InitializeBkmRequest) (entity.PaymentResponse, error) {
	response, auth, err := u.PaymentRepo.InitializeBkm(initialization)
	if err != nil {
		return entity.PaymentResponse{}, err
	}

	retrieveBkmResultDto := entity.RetrieveBkmResultRequest{
		Locale:         response.Data.Locale,
		ConversationID: response.Data.ConversationID,
		Token:          response.Data.Token,
	}

	return u.PaymentRepo.RetrieveBkmResult(retrieveBkmResultDto, auth)
}
