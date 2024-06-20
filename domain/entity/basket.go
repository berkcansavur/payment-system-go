package entity

type BasketDto struct {
	Id            string      	`json:"id"`
	Client        ClientDto     `json:"client"`
	BasketItems []BasketItem 	`json:"basketItems"`
	TotalPrice    string        `json:"totalPrice"`
	IsActive        bool 		`json:"isActive"`
}

