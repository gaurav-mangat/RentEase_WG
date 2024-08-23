// internal/domain/repositories/request_repo.go
package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
)

type RequestRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewRequestRepo initializes a new RequestRepo with a MongoDB connection.
func NewRequestRepo(uri string, dbName string, collectionName string) (interfaces.RequestRepo, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB!")

	collection := client.Database(dbName).Collection(collectionName)
	return &RequestRepo{
		client:     client,
		collection: collection,
	}, nil
}

func (repo *RequestRepo) SaveRequest(request entities.Request) error {
	_, err := repo.collection.InsertOne(context.TODO(), request)
	return err
}

func (repo *RequestRepo) FindByTenantUsername(ctx context.Context, tenantUsername string) ([]entities.Request, error) {
	filter := bson.D{{"tenantName", tenantUsername}}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var requests []entities.Request
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, err
	}

	return requests, nil
}

func (repo *RequestRepo) FindByLandlordName(ctx context.Context, landlordName string) ([]entities.Request, error) {
	filter := bson.D{{"landlordName", landlordName}}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var requests []entities.Request
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, err
	}

	return requests, nil
}

func (repo *RequestRepo) UpdateRequest(request entities.Request, status string) error {
	filter := bson.M{"_id": request.ID}
	update := bson.M{"$set": bson.M{"requestStatus": status}}
	_, err := repo.collection.UpdateOne(context.Background(), filter, update)

	return err
}
