package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"payment-system/domain/entity"
	"payment-system/pkg/usecase"
)

type PaymentController struct {
	Usecase *usecase.PaymentUsecase
}

func (c *PaymentController) ProcessBkm(w http.ResponseWriter, r *http.Request) {
	var initalization entity.InitializeBkmRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &initalization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.Usecase.ProcessBkm(initalization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
