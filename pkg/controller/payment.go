package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"payment-system/domain/entity"
	"payment-system/pkg/usecase"
)

type PaymentController struct {
	Usecase *usecase.PaymentUsecase
}

func (c *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var createPaymentRequest entity.CreatePaymentRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &createPaymentRequest)
	if err != nil {
		log.Printf("Error unmarshalling request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.Usecase.CreatePayment(createPaymentRequest)
	if err != nil {
		log.Printf("Error creating payment: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[PaymentController] CreatePayment:",response)
	json.NewEncoder(w).Encode(response)
}

