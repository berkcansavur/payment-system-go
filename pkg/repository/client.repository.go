package repository

import (
	"context"
	"fmt"
	"payment-system/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientRepository struct{
	Collection *mongo.Collection
}
func NewMongoClientRepository(client *mongo.Client, dbName, collectionName string) *ClientRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &ClientRepository{Collection: collection}
}
func (r *ClientRepository) Create(ctx context.Context, createClientDto entity.ClientDto) (*entity.ClientDto, error) {
	result, err := r.Collection.InsertOne(ctx, createClientDto)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	createClientDto.Id = insertedID

	return &createClientDto, nil
}
func (r *ClientRepository) GetById(ctx context.Context, id string) (*entity.ClientDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var client entity.ClientDto
	err = r.Collection.FindOne(ctx, filter).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}

	return &client, nil
}
func (r *ClientRepository) RemoveById(ctx context.Context, id string) (*entity.ClientDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var client entity.ClientDto
	err = r.Collection.FindOneAndDelete(ctx, filter).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}
	

	return &client, nil
}
func (r *ClientRepository) UpdateById(ctx context.Context, id string, updateClientDto entity.ClientDto) (*entity.ClientDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{
		"$set": updateClientDto,
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(1)

	var client entity.ClientDto
	err = r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}

	return &client, nil
}
func (r *ClientRepository) GetCards(ctx context.Context, id string) ([]entity.PaymentCard, error) {
	client,err := r.GetById(ctx,id)
	if err != nil {
		return nil, err
	}
	return client.Cards, nil
}
