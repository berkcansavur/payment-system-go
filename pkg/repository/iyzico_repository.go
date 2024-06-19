package repository

import (
	"os"
	"payment-system/domain/entity"

	"github.com/go-resty/resty/v2"
)

type IyzicoRepository struct{}

func (r *IyzicoRepository) InitializeBkm(initializeBkm entity.InitializeBkmRequest) (entity.PaymentResponse, entity.Authorization, error) {
	client := resty.New()
	apiURL := os.Getenv("IYZICO_BASE_URL")
	apiKey := os.Getenv("IYZICO_API_KEY")
	apiSecret := os.Getenv("IYZICO_SECRET")
	rnd := "123456789"
	//strconv.FormatInt(time.Now().UnixNano(), 10)

	authorization, pkiString := generateAuthorizationAndPkiString(apiKey, apiSecret, initializeBkm, rnd)
	auth := entity.Authorization{
		Authorization: authorization,
		PkiString:     pkiString,
	}
	resp, err := client.R().
		SetHeader("Authorization", authorization).
		SetHeader("x-iyzi-rnd", rnd).
		SetHeader("Content-Type", "application/json").
		SetHeader("pkiString", pkiString).
		SetBody(initializeBkm).
		SetResult(&entity.InitializeBkmResponse{}).
		Post(apiURL + "/payment/bkm/initialize")

	if err != nil {
		return entity.PaymentResponse{Status: "fail", Message: err.Error()}, auth, err
	}
	iyzicoResponse := resp.Result().(*entity.InitializeBkmResponse)

	if iyzicoResponse.Status == "success" {
		return entity.PaymentResponse{Status: "success", Message: "Payment successful", Data: iyzicoResponse}, auth, nil
	}
	return entity.PaymentResponse{Status: "fail", Message: iyzicoResponse.Status}, auth, nil
}
func (r *IyzicoRepository) RetrieveBkmResult(request entity.RetrieveBkmResultRequest, auth entity.Authorization) (entity.PaymentResponse, error) {
	client := resty.New()
	apiURL := os.Getenv("IYZICO_BASE_URL")
	rnd := "123456789"
	//strconv.FormatInt(time.Now().UnixNano(), 10)

	resp, err := client.R().
		SetHeader("Authorization", auth.Authorization).
		SetHeader("x-iyzi-rnd", rnd).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&entity.InitializeBkmResponse{}).
		Post(apiURL + "/payment/bkm/auth/detail")

	if err != nil {
		return entity.PaymentResponse{Status: "fail", Message: err.Error()}, err
	}
	iyzicoResponse := resp.Result().(*entity.InitializeBkmResponse)

	if iyzicoResponse.Status == "success" {
		return entity.PaymentResponse{Status: "success", Message: "Payment successful", Data: iyzicoResponse}, nil
	}
	return entity.PaymentResponse{Status: "fail", Message: iyzicoResponse.Status}, nil
}
