package main

import (
	"net/http"
	"payment-system/cmd/server"
	"payment-system/pkg/config"
	"payment-system/pkg/controller"
	"payment-system/pkg/database"
	"payment-system/pkg/repository"
	"payment-system/pkg/router"
	"payment-system/pkg/usecase"

	"github.com/gorilla/mux"
)



func main() {
	cfg := config.LoadConfig()

	// Connect to database
	client := database.Connect(cfg.DbURI)

	// Repositories
	iyzicoRepo := &repository.IyzicoRepository{}
	clientRepo := repository.NewMongoClientRepository(client, "mydatabase", "client")
	basketRepo := repository.NewMongoBasketRepository(client, "mydatabase", "basket")

	// Use cases
	paymentUsecase := &usecase.PaymentUsecase{
		PaymentRepo: iyzicoRepo,
	}
	clientUsecase := &usecase.ClientUseCase{
		ClientRepo: clientRepo,
		PaymentRepo: iyzicoRepo, 
		BasketRepo: basketRepo, 
	}
	basketUsecase := &usecase.BasketUseCase{
		BasketRepo: basketRepo,
		ClientRepo: clientRepo,
	}

	// Controllers
	paymentController := &controller.PaymentController{Usecase: paymentUsecase}
	clientController := &controller.ClientController{Usecase: clientUsecase}
	basketController := &controller.BasketController{Usecase: basketUsecase}
	// Initialize the router
	r := mux.NewRouter()

	// Register the routes
	router.ClientRouter(r, clientController)
	router.PaymentRouter(r, paymentController)
	router.BasketRouter(r, basketController)

	http.Handle("/", r)

	server.Start(r)
}
