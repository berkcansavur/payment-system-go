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
type InitializeBkmRequest struct {
	Locale              string       `json:"locale,omitempty"`
	ConversationID      string       `json:"conversationId,omitempty"`
	Price               string       `json:"price,omitempty"`
	Installment         string       `json:"installment,omitempty"`
	PaymentChannel      string       `json:"paymentChannel,omitempty"`
	BasketID            string       `json:"basketId,omitempty"`
	PaymentGroup        string       `json:"paymentGroup,omitempty"`
	PaymentCard         PaymentCard  `json:"paymentCard,omitempty"`
	Buyer               Buyer        `json:"buyer,omitempty"`
	ShippingAddress     Address      `json:"shippingAddress,omitempty"`
	BillingAddress      Address      `json:"billingAddress,omitempty"`
	BasketItems         []BasketItem `json:"basketItems,omitempty"`
	CallbackURL         string       `json:"callbackUrl,omitempty"`
	PaymentSource       string       `json:"paymentSource,omitempty"`
	Currency            string       `json:"currency,omitempty"`
	PaidPrice           string       `json:"paidPrice,omitempty"`
	ForceThreeDS        string       `json:"forceThreeDS,omitempty"`
	CardUserKey         string       `json:"cardUserKey,omitempty"`
	EnabledInstallments string       `json:"enabledInstallments,omitempty"`
}
type RetrieveBkmResultRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	Token          string `json:"token"`
}
