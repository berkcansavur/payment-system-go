package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"payment-system/domain/entity"
	"payment-system/pkg/usecase"

	"github.com/gorilla/mux"
)

type BasketController struct {
	Usecase *usecase.BasketUseCase
}

func (c *BasketController) Create(w http.ResponseWriter, r *http.Request) {
	var createBasket entity.BasketDto
	vars := mux.Vars(r)
	clientId := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &createBasket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	response, err := c.Usecase.Create(ctx, createBasket, clientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)

}