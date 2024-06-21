package entity

import "github.com/berkcansavur/iyzico-authorization"

type BasketDto struct {
	Id          string       `json:"id,omitempty" bson:"_id,omitempty"`
	Client      ClientDto    `json:"client,omitempty" bson:"client,omitempty"`
	BasketItems []iyzico.BasketItem `json:"basketItems,omitempty" bson:"basketItems,omitempty"`
	TotalPrice  string       `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	IsActive    bool         `json:"isActive,omitempty" bson:"isActive,omitempty"`
}


