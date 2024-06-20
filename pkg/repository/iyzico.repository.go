package repository

import (
	"payment-system/domain/entity"
	"payment-system/pkg/config"

	"github.com/go-resty/resty/v2"
)

type IyzicoRepository struct{}

func (r *IyzicoRepository) InitializeBkm(initializeBkm entity.InitializeBkmRequest) (entity.InitializeBkmResult, entity.Authorization, error) {
	client := resty.New()
	iyzicoConfig := config.GetIyzicoConfig()
	authorization, pkiString := generateAuthorizationAndPkiString(iyzicoConfig.APIKey, iyzicoConfig.APISecret, initializeBkm, iyzicoConfig.Rnd)

	auth := entity.Authorization{
		Authorization: authorization,
		PkiString:     pkiString,
	}
	resp, err := client.R().
		SetHeader("Authorization", authorization).
		SetHeader("x-iyzi-rnd", iyzicoConfig.Rnd).
		SetHeader("Content-Type", "application/json").
		SetHeader("pkiString", pkiString).
		SetBody(initializeBkm).
		SetResult(&entity.InitializeBkmResponse{}).
		Post(iyzicoConfig.BaseURL + "/payment/bkm/initialize")

	if err != nil {
		return entity.InitializeBkmResult{Status: "fail", Message: err.Error()}, auth, err
	}
	iyzicoResponse := resp.Result().(*entity.InitializeBkmResponse)

	if iyzicoResponse.Status == "success" {
		return entity.InitializeBkmResult{Status: "success", Message: "Payment successful", Data: iyzicoResponse}, auth, nil
	}
	return entity.InitializeBkmResult{Status: "fail", Message: iyzicoResponse.Status}, auth, err
}
func (r *IyzicoRepository) RetrieveBkmResult(resultArgs entity.RetrieveBkmResultRequest, auth entity.Authorization) (entity.InitializeBkmResult, error) {
	client := resty.New()
	iyzicoConfig := config.GetIyzicoConfig()

	resp, err := client.R().
		SetHeader("Authorization", auth.Authorization).
		SetHeader("x-iyzi-rnd", iyzicoConfig.Rnd).
		SetHeader("Content-Type", "application/json").
		SetBody(resultArgs).
		SetResult(&entity.InitializeBkmResponse{}).
		Post(iyzicoConfig.BaseURL + "/payment/bkm/auth/detail")

	if err != nil {
		return entity.InitializeBkmResult{Status: "fail", Message: err.Error()}, err
	}
	iyzicoResponse := resp.Result().(*entity.InitializeBkmResponse)

	if iyzicoResponse.Status == "success" {
		return entity.InitializeBkmResult{Status: "success", Message: "Payment successful", Data: iyzicoResponse}, nil
	}
	return entity.InitializeBkmResult{Status: "fail", Message: iyzicoResponse.Status}, nil
}
func (r *IyzicoRepository) CreatePayment(createPayment entity.CreatePaymentRequest) (entity.CreatePaymentResult, error) {
	client := resty.New()
	iyzicoConfig := config.GetIyzicoConfig()
	authorization, pkiString := generateAuthorizationAndPkiStringForCreatePayment(iyzicoConfig.APIKey, iyzicoConfig.APISecret, createPayment, iyzicoConfig.Rnd)
	resp, err := client.R().
		SetHeader("Authorization", authorization).
		SetHeader("pkiString", pkiString).
		SetHeader("x-iyzi-rnd", iyzicoConfig.Rnd).
		SetHeader("Content-Type", "application/json").
		SetBody(createPayment).
		SetResult(&entity.CreatePaymentResponse{}).
		Post(iyzicoConfig.BaseURL + "/payment/auth")

	if err != nil {
		return entity.CreatePaymentResult{Status: "fail", Message: err.Error()}, err
	}
	iyzicoResponse := resp.Result().(*entity.CreatePaymentResponse)

	if iyzicoResponse.Status == "success" {
		return entity.CreatePaymentResult{Status: "success", Message: "Payment successful", Data: iyzicoResponse}, nil
	}
	return entity.CreatePaymentResult{Status: "failure", Message: iyzicoResponse.Status}, nil
}
