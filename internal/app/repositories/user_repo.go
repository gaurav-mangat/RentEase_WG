package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"rentease/pkg/utils"
)

type UserRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewUserRepo initializes a new UserRepo with a MongoDB connection.
func NewUserRepo(uri string, dbName string, collectionName string) (interfaces.UserRepo, error) {
	client, err := connectToMongoDB(uri)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepo{
		client:     client,
		collection: collection,
	}, nil
}

// connectToMongoDB creates a new MongoDB client and connects to the database.
func connectToMongoDB(uri string) (*mongo.Client, error) {
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
	return client, nil
}

// SaveUser saves a user to the MongoDB collection.
func (repo *UserRepo) SaveUser(user entities.User) error {
	_, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	filter := bson.D{{"username", username}}

	err := repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No document found
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err) // Wrap other errors
	}

	return &user, nil
}

// CheckPassword verifies the user's password.
func (repo *UserRepo) CheckPassword(ctx context.Context, username, password string) (bool, error) {
	user, err := repo.FindByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil // User not found
	}

	// Compare the provided password with the stored hashed password
	return utils.CheckPasswordHash(password, user.PasswordHash), nil
}

func (repo *UserRepo) UpdateUser(user entities.User) error {
	filter := bson.M{"username": user.Username}
	update := bson.M{"$set": user}
	result, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

//Admin related

func (ur *UserRepo) FindAll() ([]entities.User, error) {
	ctx := context.TODO()
	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []entities.User
	err = cursor.All(ctx, &users)
	return users, err
}

func (ur *UserRepo) Delete(username string) error {
	ctx := context.TODO()
	_, err := ur.collection.DeleteOne(ctx, bson.M{"username": username})
	return err
}
