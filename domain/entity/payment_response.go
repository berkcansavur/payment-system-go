package entity

type InitializeBkmResult struct {
	Status  string
	Message string
	Fee     float64
	Data    *InitializeBkmResponse
}
type CreatePaymentResult struct {
	Status  string
	Message string
	Fee     float64
	Data    *CreatePaymentResponse
}
type PaymentResult struct {
	Status  string
	Message string
	Fee     float64
	Data    *CreatePaymentResponse
}
