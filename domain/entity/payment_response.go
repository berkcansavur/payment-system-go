package entity

type PaymentResponse struct {
	Status  string
	Message string
	Fee     float64
	Data    *InitializeBkmResponse
}
