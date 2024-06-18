package repository

import (
	"log"
	"os"
	"payment-system/domain/entity"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type IyzicoRepository struct{}

func (r *IyzicoRepository) ProcessPayment(payment entity.Payment) (entity.PaymentResponse, error) {
	client := resty.New()
	apiURL := os.Getenv("IYZICO_BASE_URL")
	apiKey := os.Getenv("IYZICO_API_KEY")
	apiSecret := os.Getenv("IYZICO_SECRET")

	amountStr := strconv.FormatFloat(payment.Amount, 'f', 2, 64)
	request := entity.IyzicoPaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          amountStr,
		PaidPrice:      amountStr,
		Currency:       "TRY",
		Installment:    1,
		BasketId:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		PaymentCard: entity.PaymentCard{
			CardHolderName: payment.CardHolderName,
			CardNumber:     payment.CardNumber,
			ExpireMonth:    payment.ExpireMonth,
			ExpireYear:     payment.ExpireYear,
			Cvc:            payment.Cvc,
			RegisterCard:   0,
		},
		Buyer: entity.Buyer{
			Id:                  "BY789",
			Name:                payment.BuyerName,
			Surname:             payment.BuyerSurname,
			IdentityNumber:      payment.IdentityNumber,
			Email:               payment.BuyerEmail,
			GsmNumber:           "+905350000000",
			RegistrationDate:    "2013-04-21 15:12:09",
			LastLoginDate:       "2015-10-05 12:43:35",
			RegistrationAddress: payment.BuyerAddress,
			City:                payment.BuyerCity,
			Country:             payment.BuyerCountry,
			ZipCode:             payment.BuyerZipCode,
			Ip:                  payment.BuyerIp,
		},
		ShippingAddress: entity.Address{
			Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			ZipCode:     "34742",
			ContactName: "Jane Doe",
			City:        "Istanbul",
			Country:     "Turkey",
		},
		BillingAddress: entity.Address{
			Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			ZipCode:     "34742",
			ContactName: "Jane Doe",
			City:        "Istanbul",
			Country:     "Turkey",
		},
		BasketItems: []entity.BasketItem{
			{
				Id:        "BI101",
				Price:     "0.3",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
			},
			{
				Id:        "BI102",
				Price:     "0.5",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
			},
			{
				Id:        "BI103",
				Price:     "0.2",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
			},
		},
		CallbackUrl: "https://www.merchant.com/callback",
	}

	// Random string for `x-iyzi-rnd` header
	rnd := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Authorization header
	authorization := apiKey + ":" + apiSecret

	resp, err := client.R().
		SetHeader("Authorization", "IYZWS "+authorization).
		SetHeader("x-iyzi-rnd", rnd).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&entity.IyzicoPaymentResponse{}).
		Post(apiURL + "/payment/bkm/initialize")

	if err != nil {
		return entity.PaymentResponse{Status: "fail", Message: err.Error()}, err
	}

	iyzicoResponse := resp.Result().(*entity.IyzicoPaymentResponse)
	log.Println(iyzicoResponse)

	if iyzicoResponse.Status == "success" {
		return entity.PaymentResponse{Status: "success", Message: "Payment successful"}, nil
	}

	return entity.PaymentResponse{Status: "fail", Message: iyzicoResponse.ErrorMessage}, nil
}
