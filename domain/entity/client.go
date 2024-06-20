package entity

type Client struct {
	Id                  string `json:"id" bson:"_id,omitempty"`
	Name                string `json:"name" bson:"name"`
	Surname             string `json:"surname" bson:"surname"`
	IdentityNumber      string `json:"identityNumber" bson:"identityNumber"`
	Email               string `json:"email" bson:"email"`
	GsmNumber           string `json:"gsmNumber" bson:"gsmNumber"`
	RegistrationDate    string `json:"registrationDate" bson:"registrationDate"`
	LastLoginDate       string `json:"lastLoginDate" bson:"lastLoginDate"`
	RegistrationAddress string `json:"registrationAddress" bson:"registrationAddress"`
	City                string `json:"city" bson:"city"`
	Country             string `json:"country" bson:"country"`
	ZipCode             string `json:"zipCode" bson:"zipCode"`
	Ip                  string `json:"ip" bson:"ip"`
	Cards				[]PaymentCard
	ShippingAddress 	Address         `json:"shippingAddress"`
	BillingAddress  	BillingAddress  `json:"billingAddress"`
}
type CreateClientDto struct {
	Id                  string 			`json:"id" bson:"_id,omitempty"`
	Name                string 			`json:"name" bson:"name"`
	Surname             string 			`json:"surname" bson:"surname"`
	IdentityNumber      string 			`json:"identityNumber" bson:"identityNumber"`
	Email               string 			`json:"email" bson:"email"`
	GsmNumber           string 			`json:"gsmNumber" bson:"gsmNumber"`
	RegistrationDate    string 			`json:"registrationDate" bson:"registrationDate"`
	LastLoginDate       string 			`json:"lastLoginDate" bson:"lastLoginDate"`
	RegistrationAddress string 			`json:"registrationAddress" bson:"registrationAddress"`
	City                string 			`json:"city" bson:"city"`
	Country             string 			`json:"country" bson:"country"`
	ZipCode             string 			`json:"zipCode" bson:"zipCode"`
	Ip                  string 			`json:"ip" bson:"ip"`
	Cards				[]PaymentCard
	ShippingAddress 	Address         `json:"shippingAddress" bson:"shippingAddress"`
	BillingAddress  	BillingAddress  `json:"billingAddress" bson:"billingAddress"`
}