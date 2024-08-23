package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// Abstract Class for User
type User struct {
	Username     string               `json:"username"`
	PasswordHash string               `json:"password_hash"`
	Name         string               `json:"name"`
	Age          int                  `json:"age"`
	Email        string               `json:"email"`
	PhoneNumber  string               `json:"phone_number"`
	Address      string               `json:"address"`
	Role         string               `json:"role"`
	Wishlist     []primitive.ObjectID `json:"wishlist"` // List of property IDs in the wishlist
}

type NormalUser struct {
	User
	Wishlist        []int `json:"wishlist"`
	PropertyListing []int `json:"propertyListing"`
}

type AdminUser struct {
	User
}
