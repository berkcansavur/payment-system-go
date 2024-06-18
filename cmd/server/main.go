package main

import (
	"log"
	"net/http"
	"os"
	"payment-system/pkg/controller"
	"payment-system/pkg/repository"
	"payment-system/pkg/usecase"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	apiKey := os.Getenv("IYZICO_API_KEY")
	apiSecret := os.Getenv("IYZICO_SECRET")

	log.Println("API Key:", apiKey)
	log.Println("API Secret:", apiSecret)

	iyzicoRepo := &repository.IyzicoRepository{}
	paymentUsecase := &usecase.PaymentUsecase{PaymentRepo: iyzicoRepo}
	paymentController := &controller.PaymentController{Usecase: paymentUsecase}

	http.HandleFunc("/payment", paymentController.ProcessPayment)
	http.ListenAndServe(":8080", nil)
}
