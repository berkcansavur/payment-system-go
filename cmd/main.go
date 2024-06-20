package main

import (
	"log"
	"net/http"
	"payment-system/pkg/config"
	"payment-system/pkg/controller"
	"payment-system/pkg/database"
	"payment-system/pkg/repository"
	"payment-system/pkg/router"
	"payment-system/pkg/usecase"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
)


func main() {
	cfg := config.LoadConfig()
	log.Println("API Key:", cfg.ApiKey)
	log.Println("API Secret:", cfg.ApiSecret)

	// Connect to database
	client := database.Connect(cfg.DbURI)

	// Repositories
	iyzicoRepo := &repository.IyzicoRepository{}
	clientRepo := repository.NewMongoClientRepository(client, "mydatabase", "client")

	// Use cases
	paymentUsecase := &usecase.PaymentUsecase{PaymentRepo: iyzicoRepo}
	clientUsecase := &usecase.ClientUseCase{ClientRepo: clientRepo}

	// Controllers
	paymentController := &controller.PaymentController{Usecase: paymentUsecase}
	clientController := &controller.ClientController{Usecase: clientUsecase}

	// Initialize the router
	r := mux.NewRouter()

	// Register the routes
	router.ClientRouter(r, clientController)
	router.PaymentRouter(r, paymentController)

	http.Handle("/", r)

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
