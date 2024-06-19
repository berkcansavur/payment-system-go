package repository

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"payment-system/domain/entity"
	"payment-system/pkg/formatter"
)

// Generates authorization string and PKI string
func generateAuthorizationAndPkiString(apiKey string, secretKey string, request entity.InitializeBkmRequest, rnd string) (string, string) {
	requestString := formatter.FormatInitializeBkm(request)
	hash := sha1.New()
	hash.Write([]byte(apiKey + rnd + secretKey + requestString))
	hashInBase64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	authorization := fmt.Sprintf("IYZWS %s:%s", apiKey, hashInBase64)
	pkiString := apiKey + rnd + secretKey + requestString
	return authorization, pkiString
}
