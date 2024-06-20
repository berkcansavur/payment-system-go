package router

import (
	"payment-system/pkg/controller"

	"github.com/gorilla/mux"
)

func ClientRouter(r *mux.Router, clientController *controller.ClientController) {
	// Client routes
	r.HandleFunc("/client/create", clientController.Register).Methods("POST")
	r.HandleFunc("/client/{id}", clientController.Delete).Methods("DELETE")
	r.HandleFunc("/client/{id}", clientController.Update).Methods("POST")
	r.HandleFunc("/client/{id}", clientController.Get).Methods("GET")
	r.HandleFunc("/client/{id}/cards", clientController.GetCards).Methods("GET")
	r.HandleFunc("/client/{id}/cards", clientController.AddCard).Methods("POST")
	r.HandleFunc("/client/{id}/cards", clientController.RemoveCard).Methods("DELETE")
}

func PaymentRouter(r *mux.Router, paymentController *controller.PaymentController) {
	// Payment routes
	r.HandleFunc("/payment/bkm", paymentController.ProcessBkm).Methods("POST")
	r.HandleFunc("/payment/create", paymentController.CreatePayment).Methods("POST")
}
func BasketRouter(r *mux.Router, basketController *controller.BasketController){
		// Basket routes
	r.HandleFunc("/basket/create", basketController.Create).Methods("POST")
}