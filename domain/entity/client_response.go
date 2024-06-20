package entity
type CreateClientResponse struct {
	Status  			string			`json:"status" bson:"status,omitempty"`
	Data                *ClientDto			`json:"data"`
}