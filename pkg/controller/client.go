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

type ClientController struct {
	Usecase *usecase.ClientUseCase
}

func (c *ClientController) Register(w http.ResponseWriter, r *http.Request) {
	var createClient entity.ClientDto
	body, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &createClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := c.Usecase.Register(ctx, createClient)
	if err != nil {
		log.Printf("Error registering client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] Register:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	response, err := c.Usecase.Delete(ctx, id)
	if err != nil {
		log.Printf("Error deleting client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] Delete:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) Update(w http.ResponseWriter, r *http.Request) {
	var updateClient entity.ClientDto
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &updateClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := c.Usecase.Update(ctx, id, updateClient)
	if err != nil {
		log.Printf("Error updating client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] Update:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	response, err := c.Usecase.Get(ctx, id)
	if err != nil {
		log.Printf("Error getting client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] Get:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) GetCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	response, err := c.Usecase.GetClientsCards(ctx, id)
	if err != nil {
		log.Printf("Error getting client cards: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] GetCards:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) AddCard(w http.ResponseWriter, r *http.Request) {
	var card entity.PaymentCard
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := vars["id"]
	ctx := r.Context()
	response, err := c.Usecase.AddCard(ctx, id, card)
	if err != nil {
		log.Printf("Error adding card to client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] AddCard:",response)
	json.NewEncoder(w).Encode(response)
}

func (c *ClientController) RemoveCard(w http.ResponseWriter, r *http.Request) {
	var card entity.PaymentCard
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := c.Usecase.RemoveCard(ctx, id, card)
	if err != nil {
		log.Printf("Error removing card from client: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] RemoveCard:",response)
	json.NewEncoder(w).Encode(response)
}
func (c *ClientController) Checkout(w http.ResponseWriter, r *http.Request) {
	var checkoutDto entity.CheckoutDto
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &checkoutDto)
	if err != nil {
		log.Printf("Error unmarshalling request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response, err := c.Usecase.CheckoutClientsActiveBasket(ctx, id, checkoutDto)
	if err != nil {
		log.Printf("Error checking out client's active basket: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[ClientController] Checkout:",response)
	json.NewEncoder(w).Encode(response)
}
