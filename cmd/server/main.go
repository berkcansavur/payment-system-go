package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"payment-system/pkg/controller"
	"payment-system/pkg/repository"
	"payment-system/pkg/usecase"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

func main() {
	apiKey := os.Getenv("IYZICO_API_KEY")
	apiSecret := os.Getenv("IYZICO_SECRET")

	log.Println("API Key:", apiKey)
	log.Println("API Secret:", apiSecret)

	iyzicoRepo := &repository.IyzicoRepository{}
	paymentUsecase := &usecase.PaymentUsecase{PaymentRepo: iyzicoRepo}
	paymentController := &controller.PaymentController{Usecase: paymentUsecase}
	
	http.HandleFunc("/payment/bkm", paymentController.ProcessBkm)
	http.HandleFunc("/payment/create", paymentController.CreatePayment)

	clientRepo := repository.NewMongoClientRepository(client, "mydatabase", "client") 
	clientUsecase := &usecase.ClientUseCase{ClientRepo: clientRepo}
	clientController := &controller.ClientController{Usecase: clientUsecase}
	r := mux.NewRouter()
	r.HandleFunc("/client/create", clientController.Register).Methods("POST")
	r.HandleFunc("/client/{id}", clientController.Delete).Methods("DELETE")
	r.HandleFunc("/client/{id}", clientController.Update).Methods("POST")
	r.HandleFunc("/client/{id}", clientController.Get).Methods("GET")
	r.HandleFunc("/client/cards/{id}", clientController.GetCards).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
