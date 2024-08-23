package services

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"rentease/pkg/utils"
)

type UserService struct {
	userRepo interfaces.UserRepo
}

func NewUserService(userRepo interfaces.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

//func (us *UserService) Login(user entities.User) error {
//}

func (us *UserService) SignUp(user entities.User) bool {
	err := us.userRepo.SaveUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return true

}

func (us *UserService) FindByUsername(username string) (entities.User, error) {
	user, err := us.userRepo.FindByUsername(context.Background(), username)
	if err != nil {
		return entities.User{}, err
	}
	if user == nil {
		return entities.User{}, nil // No user found
	}
	return *user, nil
}

func (us *UserService) Login(username, password string) (bool, error) {
	exist, err := us.userRepo.CheckPassword(context.Background(), username, password)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	utils.ActiveUser = username
	return true, nil

}

func (us *UserService) AddToWishlist(username string, propertyID primitive.ObjectID) error {
	ctx := context.TODO() // Use a proper context in real applications

	user, err := us.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	// Check if the property is already in the wishlist
	for _, id := range user.Wishlist {
		if id == propertyID {
			return errors.New("property is already in the wishlist")
		}
	}

	// Add the property ID to the wishlist
	user.Wishlist = append(user.Wishlist, propertyID)

	// Update the user record
	err = us.userRepo.UpdateUser(*user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(user entities.User) error {
	return us.userRepo.UpdateUser(user)
}

// Admin specific services
func (us *UserService) GetAllUsers() ([]entities.User, error) {
	return us.userRepo.FindAll()
}

func (us *UserService) DeleteUser(username string) error {
	return us.userRepo.Delete(username)
}
