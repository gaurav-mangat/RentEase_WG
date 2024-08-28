package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type UserService interface {
	SignUp(user entities.User) error
	FindByUsername(username string) (entities.User, error)
	Login(username, password string) (bool, error)
	AddToWishlist(username string, propertyID primitive.ObjectID) error
	UpdateUser(user entities.User) error
	GetAllUsers() ([]entities.User, error)
	DeleteUser(username string) error
}
