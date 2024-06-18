package entity

type IyzicoPaymentRequest struct {
	Locale          string       `json:"locale"`
	ConversationId  string       `json:"conversationId"`
	Price           string       `json:"price"`
	PaidPrice       string       `json:"paidPrice"`
	Currency        string       `json:"currency"`
	Installment     int          `json:"installment"`
	BasketId        string       `json:"basketId"`
	PaymentChannel  string       `json:"paymentChannel"`
	PaymentGroup    string       `json:"paymentGroup"`
	PaymentCard     PaymentCard  `json:"paymentCard"`
	Buyer           Buyer        `json:"buyer"`
	ShippingAddress Address      `json:"shippingAddress"`
	BillingAddress  Address      `json:"billingAddress"`
	BasketItems     []BasketItem `json:"basketItems"`
	CallbackUrl     string       `json:"callbackUrl"`
}
type PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireMonth    string `json:"expireMonth"`
	ExpireYear     string `json:"expireYear"`
	Cvc            string `json:"cvc"`
	RegisterCard   int    `json:"registerCard"`
}
type Buyer struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber"`
	RegistrationDate    string `json:"registrationDate"`
	LastLoginDate       string `json:"lastLoginDate"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zipCode"`
	Ip                  string `json:"ip"`
}
type Address struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
type BasketItem struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2"`
	ItemType  string `json:"itemType"`
	Price     string `json:"price"`
}
