package repository

import (
	"log"
	"os"
	"payment-system/domain/entity"
	"strconv"

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
    }

    request.PaymentCard.CardHolderName = payment.CardHolderName
    request.PaymentCard.CardNumber = payment.CardNumber
    request.PaymentCard.ExpireMonth = payment.ExpireMonth
    request.PaymentCard.ExpireYear = payment.ExpireYear
    request.PaymentCard.Cvc = payment.Cvc
    request.PaymentCard.RegisterCard = 0

    request.Buyer = struct {
        Id                  string `json:"id"`
        Name                string `json:"name"`
        Surname             string `json:"surname"`
        IdentityNumber      string `json:"identityNumber"`
        Email               string `json:"email"`
        GsmNumber           string `json:"gsmNumber"`
        RegistrationDate    string `json:"registrationDate"`
        LastLoginDate       string `json:"lastLoginDate"`
        RegistrationAddress string `json:"registrationAddress"`
        City                string `json:"city"`
        Country             string `json:"country"`
        ZipCode             string `json:"zipCode"`
        Ip                  string `json:"ip"`
    }{
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
    }

    resp, err := client.R().
        SetBasicAuth(apiKey, apiSecret).
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
