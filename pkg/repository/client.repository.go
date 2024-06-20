package repository

import (
	"context"
	"payment-system/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct{
	Collection *mongo.Collection
}
func NewMongoClientRepository(client *mongo.Client, dbName, collectionName string) *ClientRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &ClientRepository{Collection: collection}
}
func (r *ClientRepository) Create(ctx context.Context, createClientDto entity.CreateClientDto) (entity.CreateClientResponse,error){
	result, err := r.Collection.InsertOne(ctx, createClientDto)
	if err!= nil {
        return entity.CreateClientResponse{}, err
    }
	id := result.InsertedID.(primitive.ObjectID).Hex()
	client := &entity.Client{Id: id}
	createClientResponse := entity.CreateClientResponse{
		Status: "success",
		Data:   client,
	}
	return createClientResponse, nil
}