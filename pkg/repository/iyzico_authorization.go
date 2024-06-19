package repository

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"payment-system/domain/entity"
	"payment-system/pkg/formatter"
)

func generateAuthorizationAndPkiString(apiKey string, secretKey string, request entity.InitializeBkmRequest, rnd string) (string, string) {
	requestString := formatter.FormatInitializeBkm(request)
	return getAuthorizationAndPkiString(apiKey, rnd, secretKey, requestString)
}
func generateAuthorizationAndPkiStringForCreatePayment(apiKey string, secretKey string, request entity.CreatePaymentRequest, rnd string) (string, string) {
	requestString := formatter.FormatCreatePayment(request)
	return getAuthorizationAndPkiString(apiKey, rnd, secretKey, requestString)
}
func getAuthorizationAndPkiString(apiKey string, rnd string, secretKey string, requestString string) (string, string) {
	hash := sha1.New()
	hash.Write([]byte(apiKey + rnd + secretKey + requestString))
	hashInBase64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	authorization := fmt.Sprintf("IYZWS %s:%s", apiKey, hashInBase64)
	pkiString := apiKey + rnd + secretKey + requestString
	return authorization, pkiString
}
