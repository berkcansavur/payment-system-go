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

type BasketRepository struct{
	Collection *mongo.Collection
}
func NewMongoBasketRepository(client *mongo.Client, dbName, collectionName string) *BasketRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &BasketRepository{Collection: collection}
}
func (r *BasketRepository) Create(ctx context.Context, createBasketDto entity.BasketDto) (*entity.BasketDto, error) {
	result, err := r.Collection.InsertOne(ctx, createBasketDto)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	createBasketDto.Id = insertedID

	return &createBasketDto, nil
}
func (r *BasketRepository) GetById(ctx context.Context, id string) (*entity.BasketDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var client entity.BasketDto
	err = r.Collection.FindOne(ctx, filter).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}

	return &client, nil
}
func (r *BasketRepository) RemoveById(ctx context.Context, id string) (*entity.BasketDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var client entity.BasketDto
	err = r.Collection.FindOneAndDelete(ctx, filter).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}
	

	return &client, nil
}
func (r *BasketRepository) UpdateById(ctx context.Context, id string, updateBasketDto entity.BasketDto) (*entity.BasketDto, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %v", err)
	}
	
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{
		"$set": bson.M{
			"client": 		updateBasketDto.Client,
			"basketItems":	updateBasketDto.BasketItems,
			"totalPrice":	updateBasketDto.TotalPrice,
			"isActive":		updateBasketDto.IsActive,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(1)

	var client entity.BasketDto
	err = r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with id %s", id)
		}
		return nil, fmt.Errorf("error finding document with id %s: %v", id, err)
	}

	return &client, nil
}
func (r *BasketRepository) GetActiveBasketByClientId(ctx context.Context, clientId string) (*entity.BasketDto, error) {
	objID, err := primitive.ObjectIDFromHex(clientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client id format: %v", err)
	}
	
	filter := bson.M{"client._id": objID, "isActive": true}

	var basket entity.BasketDto
	err = r.Collection.FindOne(ctx, filter).Decode(&basket)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no active basket found for client id %s", clientId)
		}
		return nil, fmt.Errorf("error finding active basket for client id %s: %v", clientId, err)
	}

	return &basket, nil
}

