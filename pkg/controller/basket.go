package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	clientID := vars["clientId"]
	body, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()
	if err != nil {
		log.Printf("Error reading request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &createBasket)
	if err != nil {
		log.Printf("Error unmarshalling request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response, err := c.Usecase.Create(ctx, createBasket, clientID)
	if err != nil {
		log.Printf("Error creating basket: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[BasketController] Create:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *BasketController) Abort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()

	response, err := c.Usecase.Abort(ctx, id)
	if err != nil {
		log.Printf("Error aborting basket: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[BasketController] Abort:",response)
	json.NewEncoder(w).Encode(response)
}
