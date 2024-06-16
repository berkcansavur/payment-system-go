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

func (c *PaymentController) ProcessPayment(w http.ResponseWriter, r *http.Request) {
    var payment entity.Payment
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = json.Unmarshal(body, &payment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response, err := c.Usecase.Execute(payment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(response)
}
