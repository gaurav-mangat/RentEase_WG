package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username     string               `bson:"username"`
	PasswordHash string               `bson:"password_hash"`
	Name         string               `bson:"name"`
	Age          int                  `bson:"age"`
	Email        string               `bson:"email"`
	PhoneNumber  string               `bson:"phone_number"`
	Address      string               `bson:"address"`
	Role         string               `bson:"role"`
	Wishlist     []primitive.ObjectID `bson:"wishlist"` // List of property IDs in the wishlist
}
