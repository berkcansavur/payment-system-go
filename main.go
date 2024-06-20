package main

import (
	"log"
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
	log.Println("Starting the payment system application...")

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	client := database.Connect(cfg.DbURI)

	// Initialize repositories
	iyzicoRepo := &repository.IyzicoRepository{}
	clientRepo := repository.NewMongoClientRepository(client, "mydatabase", "client")
	basketRepo := repository.NewMongoBasketRepository(client, "mydatabase", "basket")
	log.Println("Repositories initialized successfully.")

	// Initialize use cases
	paymentUsecase := &usecase.PaymentUsecase{
		PaymentRepo: iyzicoRepo,
	}
	clientUsecase := &usecase.ClientUseCase{
		ClientRepo:  clientRepo,
		PaymentRepo: iyzicoRepo,
		BasketRepo:  basketRepo,
	}
	basketUsecase := &usecase.BasketUseCase{
		BasketRepo: basketRepo,
		ClientRepo: clientRepo,
	}
	log.Println("Use cases initialized successfully.")

	// Initialize controllers
	paymentController := &controller.PaymentController{Usecase: paymentUsecase}
	clientController := &controller.ClientController{Usecase: clientUsecase}
	basketController := &controller.BasketController{Usecase: basketUsecase}
	log.Println("Controllers initialized successfully.")

	// Initialize the router
	r := mux.NewRouter()

	// Register the routes
	router.ClientRouter(r, clientController)
	router.PaymentRouter(r, paymentController)
	router.BasketRouter(r, basketController)
	log.Println("Routes registered successfully.")

	// Start the server
	log.Println("Starting the server...")
	server.Start(r)
}
